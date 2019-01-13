package helper

import (
	"strconv"

	ceria "github.com/zokypesch/ceria"
	"github.com/zokypesch/ceria/core"
	"github.com/zokypesch/ceria/helper"
	hlp "github.com/zokypesch/ceria/helper"
	repo "github.com/zokypesch/ceria/repository"
	routeService "github.com/zokypesch/ceria/route"
)

// GetDB for get db service wrapping
func GetDB() *core.Connection {
	config := helper.NewReadConfigService()
	config.Init()

	port, _ := strconv.Atoi(config.GetByName("db.port"))

	db := core.NewServiceConnection(
		config.GetByName("db.driver"),
		config.GetByName("db.host"),
		port,
		config.GetByName("db.user"),
		config.GetByName("db.password"),
		config.GetByName("db.name"),
	)
	return db
}

// GetRouter func for get router
func GetRouter() (*routeService.GinCfg, *repo.ElasticProperties) {
	config := helper.NewReadConfigService()
	config.Init()

	initRouter := routeService.NewRouteService(true, "./templates", false)

	myElastic := &repo.ElasticProperties{}
	if config.GetByName("elastic.status") == "true" {
		myElastic = &repo.ElasticProperties{
			Status: true,
			Host:   config.GetByName("elastic.host"),
			Port:   config.GetByName("elastic.port"),
		}
	}

	return initRouter, myElastic
}

// GetGroup for get group default nil
func GetGroup(name string) *ceria.GroupConfiguration {
	return &ceria.GroupConfiguration{
		Name:       name,
		Middleware: nil,
	}
}

// GetRabbitMQ for get configuration in rabbitMQ
func GetRabbitMQ() (*core.RabbitMQCore, error) {
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

	return rb, errNew
}
