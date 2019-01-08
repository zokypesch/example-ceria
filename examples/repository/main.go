package main

import (
	"github.com/jinzhu/gorm"
	"github.com/zokypesch/ceria/helper"
	repo "github.com/zokypesch/ceria/repository"
	routeService "github.com/zokypesch/ceria/route"
	hlp "github.com/zokypesch/example-ceria/helper"
)

// Example for example struct repository
type Example struct {
	gorm.Model
	Title  string `validate:"required" form:"title" json:"title" binding:"required"`
	Author string `validate:"required" form:"author" json:"author" binding:"required"`
}

func main() {
	conn := hlp.GetDB()
	db, err := conn.GetConn()

	if err != nil {
		panic(err)
	}
	defer db.Close()
	initRouter := routeService.NewRouteService(true, "../../templates/*", true)

	r, errRouting := initRouter.Register(true)

	if errRouting != nil {
		panic(errRouting)
	}

	config := helper.NewReadConfigService()
	config.Init()
	var withElastic *repo.ElasticProperties
	withElastic = &repo.ElasticProperties{}

	confStatus := config.GetByName("elastic.status")
	if confStatus == "true" {
		// elastic configuration
		withElastic = &repo.ElasticProperties{
			Status: true,
			Host:   config.GetByName("elastic.host"),
			Port:   config.GetByName("elastic.port"),
		}
	}

	myStruct := Example{}
	repos := repo.NewMasterRepository(&myStruct, db, withElastic)

	handl := repo.NewServiceRouteHandler(r, repos, &repo.QueryProps{WithPagination: true})

	handl.PathRegister()

}
