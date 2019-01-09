package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zokypesch/ceria"
	"github.com/zokypesch/ceria/core"
	hlp "github.com/zokypesch/ceria/helper"
	"github.com/zokypesch/ceria/repository"
	cr "github.com/zokypesch/example-ceria/core"
	"github.com/zokypesch/example-ceria/helper"
	"github.com/zokypesch/example-ceria/migration"
	"github.com/zokypesch/example-ceria/model"
)

func main() {
	db := helper.GetDB()
	initRouter, elastic := helper.GetRouter()
	grp := helper.GetGroup("")

	dbs, _ := db.GetConn()

	migration.NewListMigration().Migrate(dbs)

	_, errM1 := ceria.RegisterModel(
		initRouter,
		db,
		elastic,
		&model.Article{},
		grp,
		&repository.QueryProps{PreloadStatus: true, Preload: []string{"Comments", "Author"},
			WithPagination: true},
		nil,
	)

	if errM1 != nil {
		panic(errM1)
	}

	ceria.RegisterModel(
		initRouter,
		db,
		elastic,
		&model.Author{},
		grp,
		&repository.QueryProps{WithPagination: true},
		nil,
	)

	r, _ := initRouter.Register(true)

	middle := cr.NewServiceMiddleWare()
	midSet := middle.GetMiddleWareAPI()

	r.POST("/login", midSet.LoginHandler)
	r.GET("/refresh_token", midSet.RefreshHandler)

	ceria.RegisterModel(initRouter, db, elastic, &model.Comment{},
		&ceria.GroupConfiguration{Name: "auth", Middleware: midSet.MiddlewareFunc()},
		&repository.QueryProps{}, []string{"bulkcreate", "delete", "bulkdelete", "find"})

	r.GET("/newtask", pushToRabbit)
	r.Run(":9090")
}

func pushToRabbit(ctx *gin.Context) {

	msg := ctx.DefaultQuery("msg", "hello my name is udin !")

	config := hlp.NewReadConfigService()
	config.Init()

	host := config.GetByName("rabbitmq.host")
	hostname := config.GetByName("rabbitmq.hostname")
	port := config.GetByName("rabbitmq.port")
	user := config.GetByName("rabbitmq.user")
	password := config.GetByName("rabbitmq.password")

	rb, errNew := core.NewServiceRabbitMQ(&core.RabbitMQConfig{
		Host:       host,
		Hostname:   hostname,
		Port:       port,
		User:       user,
		Password:   password,
		WorkerName: "my_task",
	})

	if errNew != nil {
		hlp.NewServiceHTTPHelper().EchoResponseBadRequest(ctx, "failed to create new task", errNew.Error())
	}

	rb.RegisterNewTask(msg)

	hlp.NewServiceHTTPHelper().EchoResponseSuccess(ctx, nil)
}
