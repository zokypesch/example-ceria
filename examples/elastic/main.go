package main

import (
	"fmt"

	"github.com/zokypesch/example-ceria/model"

	core "github.com/zokypesch/ceria/core"
	"github.com/zokypesch/ceria/helper"
)

// Example of sample elastic
type Example struct {
	Name string
	Age  string
}

func main() {
	config := helper.NewReadConfigService()
	config.Init()

	host := config.GetByName("elastic.host")
	port := config.GetByName("elastic.port")
	stat := config.GetByName("elastic.status")

	if stat == "false" {
		fmt.Println("Disable elastic")
		return
	}

	var err error

	hostElastic := "http://" + host + ":" + port
	elastic, errTest := core.NewServiceElasticCore(&Example{}, hostElastic)

	if errTest != nil {
		panic(errTest)
	}

	err = elastic.DeleteIndex() // for delete all index by name examples

	if err != nil {
		panic(err)
	}

	err = elastic.AddDocument("1", &Example{"udin", "40"}) // Inserting a data

	if err != nil {
		panic(err)
	}

	err = elastic.EditDocument("1", &Example{"amril", "40"}) // Update data

	if err != nil {
		panic(err)
	}

	err = elastic.DeleteDocument("1") // Delete data
	if err != nil {
		panic(err)
	}

	elastic, _ = core.NewServiceElasticCore(&model.Article{}, hostElastic)
	elastic.DeleteIndex()

	elastic, _ = core.NewServiceElasticCore(&model.Author{}, hostElastic)
	elastic.DeleteIndex()

	elastic, _ = core.NewServiceElasticCore(&model.Comment{}, hostElastic)
	elastic.DeleteIndex()

}
