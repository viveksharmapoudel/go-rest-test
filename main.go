package main

import(
	"rest-api-test/models"
	"rest-api-test/routes"
	"github.com/jinzhu/gorm"
	"fmt"
	"rest-api-test/config"
)

var err error

func main(){
	// fmt.Println("test")
	// fmt.Println(controller.UserController())

	//configuration of db

	config.DB, err = gorm.Open("mysql", config.DbURL(config.BuildDBConfig()))

	//checking the errors
	if err != nil {
		fmt.Println("Status:", err)
	}

	//closing the database
	defer config.DB.Close()

	//migrating the setup
	config.DB.AutoMigrate(&models.Article{})

	r := routes.SetupRouter()
	//running
	r.Run()
}