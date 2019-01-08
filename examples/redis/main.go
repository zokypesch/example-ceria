package main

import (
	"fmt"

	core "github.com/zokypesch/ceria/core"
	"github.com/zokypesch/ceria/helper"
)

func main() {
	config := helper.NewReadConfigService()
	config.Init()

	host := config.GetByName("redis.host")
	port := config.GetByName("redis.port")
	stat := config.GetByName("redis.status")

	if stat == "false" {
		fmt.Println("Disable redis")
		return
	}

	cmd, err := core.NewServiceRedisCore(host, port)
	if err != nil {
		panic(err)
	}

	errCreate := cmd.CreateOrUpdateDocument("album", "1", "title", "welcome to the jungle")

	if errCreate != nil {
		panic(errCreate)
	}

	errDelete := cmd.DeleteDocument("album", "1")
	if errDelete != nil {
		panic(errDelete)
	}

	res, err := cmd.GetDocument("album", "1", "title")
	if err == nil {
		fmt.Println(res)
	}

	resMap, errG := cmd.GetAllDocument("album", "1")
	if errG == nil {
		fmt.Println(resMap)
	}
}
