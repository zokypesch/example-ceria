package main

import (
	"fmt"

	"github.com/zokypesch/ceria/helper"
)

func main() {
	help := helper.NewReadConfigService()
	err := help.Init()

	if err != nil {
		panic(err)
	}

	fmt.Println(help.GetByName("db.host"))
}
