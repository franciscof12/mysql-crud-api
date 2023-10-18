package main

import (
	"log"

	database "github.com/franciscof12/v1/crud_api/db"
	"github.com/franciscof12/v1/crud_api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	r := gin.Default()
	routes.SetUpRoutes(r, db)

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
