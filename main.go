package main

import (
	"example.com/car-rental/db"
	"example.com/car-rental/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDb()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")
}

