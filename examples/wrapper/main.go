package main

import (
	"github.com/zokypesch/ceria"
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
		&repository.QueryProps{PreloadStatus: true, Preload: []string{"Comments"}, WithPagination: true},
		nil,
	)

	if errM1 != nil {
		panic(errM1)
	}

	r, _ := initRouter.Register(true)

	middle := cr.NewServiceMiddleWare()
	midSet := middle.GetMiddleWareAPI()

	r.POST("/login", midSet.LoginHandler)
	r.GET("/refresh_token", midSet.RefreshHandler)

	ceria.RegisterModel(initRouter, db, elastic, &model.Comment{},
		&ceria.GroupConfiguration{Name: "auth", Middleware: midSet.MiddlewareFunc()},
		&repository.QueryProps{}, []string{"bulkcreate", "delete", "bulkdelete", "find"})

	r.Run(":9090")
}
