package handler

import (
	"net/http"

	"example.com/car-rental/db"
	"example.com/car-rental/middlewares"
	"example.com/car-rental/routes"
	"github.com/gin-gonic/gin"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(middlewares.CORSMiddleware())
	routes.RegisterRoutes(router)
	router.Run(":5000")
	// Call db.InitDb() to ensure the database is connected
	if db.DB == nil {
		db.InitDb() // เชื่อมต่อกับฐานข้อมูล
	}

	// Serve the request
	router.ServeHTTP(w, r)
}

func main() {
	db.InitDb()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Use(middlewares.CORSMiddleware())
	server.Run(":5000")
}

