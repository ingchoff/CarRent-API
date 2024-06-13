package routes

import (
	"example.com/car-rental/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.GET("/user", getUser)
	// authenticated.POST("/user/refreshtoken")
	authenticated.POST("/car/new", createCar)
	authenticated.GET("/cars", getCars)
	authenticated.GET("/cars/:id", getCar)
	authenticated.PUT("/car/:id", updateCar)
	authenticated.DELETE("/car/:id", deleteCar)
	authenticated.POST("/rental/new", createRental)
	authenticated.GET("/rentals", getRentals)
	authenticated.GET("/rental/:id", getRental)
	authenticated.PUT("/rental/:id", updateRental)
	authenticated.DELETE("/rental/:id", deleteRental)

	server.POST("/signup", signup)
	server.POST("/login", login)
}