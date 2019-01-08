package main

import (
	"strconv"

	"github.com/zokypesch/ceria/core"
	"github.com/zokypesch/ceria/helper"
)

func main() {
	config := helper.NewReadConfigService()
	config.Init()

	host := config.GetByName("db.host")
	port := config.GetByName("db.port")
	driver := config.GetByName("db.driver")
	user := config.GetByName("db.user")
	password := config.GetByName("db.password")
	dbname := config.GetByName("db.name")
	stat := config.GetByName("db.status")
	newPort, _ := strconv.Atoi(port)

	if stat == "false" {
		return
	}

	conn := core.NewServiceConnection(
		driver,
		host,
		newPort,
		user,
		password,
		dbname,
	)

	dbs, err := conn.GetConn()
	if err != nil {
		panic(err)
	}

	defer dbs.Close()

	// get fake connection of gorm
	db := core.GetTestConnection()

	db.LogMode(true)
	defer db.Close()

	// please read a documentation github.com/selvatico/go-mocket
}
