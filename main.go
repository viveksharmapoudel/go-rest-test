package main

import (
	"rest-api-test/models"
	"rest-api-test/routes"
	"rest-api-test/config"
	"github.com/jinzhu/gorm"
	"fmt"
)

var err error

func main()  {

	//db configuration
	config.DB, err = gorm.Open("mysql", config.DbURL(config.BuildDBConfig()))

	//error handling
	if err != nil {
		fmt.Println("Status:", err)
	}

	//db close
	defer config.DB.Close()

	//model migrate 
	config.DB.AutoMigrate(&models.Article{})

	//all the routers
	r := routes.SetupRouter()
	//running
	r.Run()

}