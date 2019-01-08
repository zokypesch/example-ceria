package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/zokypesch/ceria/core"
	"github.com/zokypesch/ceria/helper"

	mig "github.com/zokypesch/example-ceria/migration"

	"github.com/jinzhu/gorm"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	h := helper.NewReadConfigService()

	errCfg := h.Init()
	if errCfg != nil {
		panic(errCfg)
	}

	port, _ := strconv.Atoi(h.GetByName("db.port"))

	serv := core.NewServiceConnection(
		h.GetByName("db.driver"),
		h.GetByName("db.host"),
		port,
		h.GetByName("db.user"),
		h.GetByName("db.password"),
		h.GetByName("db.name"),
	)

	db, err := serv.GetConn()

	sv := InitialMenu(db)
	if err != nil {
		panic(err)
	}

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "action",
			Value: "migrate",
			Usage: "for create your table",
		},
	}

	app.Action = func(c *cli.Context) error {
		if len(c.Args()) == 0 {
			fmt.Println("please type your args")
			return nil
		}

		argsWithoutProg := c.Args()
		for _, v := range argsWithoutProg {
			if strings.Contains(v, "-") {
				continue
			}
			action := fmt.Sprintf("%s:%s", v, c.String("action"))
			sv.Menu(action)
		}
		return nil
	}

	errOS := app.Run(os.Args)
	if errOS != nil {
		log.Fatal(err)
	}
}

// MenuInterface for mocking the Menu
type MenuInterface interface {
	Menu(name string)
	Contains(value string) bool
}

// MenuStruct for list struct menu
type MenuStruct struct {
	listMenu []string
	db       *gorm.DB
	MenuInterface
}

// InitialMenu for set list menu allowed
func InitialMenu(db *gorm.DB) *MenuStruct {
	return &MenuStruct{
		listMenu: []string{
			"db:migrate",
		},
		db: db,
	}
}

// Contains for check is available for menu
func (st *MenuStruct) Contains(value string) bool {
	for _, v := range st.listMenu {
		if v == value {
			return true
		}
	}
	return false
}

// Menu for menu cli
func (st *MenuStruct) Menu(name string) {
	if !st.Contains(name) {
		fmt.Println("command list notfound !")
	}
	switch name {
	case "db:migrate":
		migrate := mig.NewListMigration()
		migrate.Migrate(st.db)
		fmt.Println("migrate success !")
		defer st.db.Close()
	}
}
