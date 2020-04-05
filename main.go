package main

import (
	"fmt"
	"gintodoapp/Config"
	"gintodoapp/Models"
	"gintodoapp/Routes"

	"github.com/jinzhu/gorm"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbUrl(Config.BuidDbConfig()))
	if err != nil {
		fmt.Println("Status :", err)
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.Todo{})

	r := Routes.SetupRouter()
	r.Run()
}
