package main

import (
	"example.com/car-rental/db"
	"example.com/car-rental/middlewares"
	"example.com/car-rental/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDb()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Use(middlewares.CORSMiddleware())
	server.Run(":5000")
}

