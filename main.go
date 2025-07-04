package main

import (
	"log"

	"github.com/api-skeleton/config"
	"github.com/api-skeleton/routes"
)

func main() {
	//connect DB
	db := config.Connect()
	defer db.Close()

	// migrate DB
	config.MigrateDB()

	// routes controller
	controller := routes.Controller()

	//log
	// fmt.Println("Application Running in Port : 8080")
	log.Fatal(controller.Run(":8080"))
}
