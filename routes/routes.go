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
	// authenticated.DELETE("/events/:id", deleteEvent)
	// authenticated.POST("/events/:id/register", registerForEvent)
	// authenticated.DELETE("/events/:id/register", cancelRegistration)
	// server.GET("/events", getEvents)    // GET, POST, PUT, PATCH, DELETE
	// server.GET("/events/:id", getEvent) // /events/1, /events/5

	server.POST("/signup", signup)
	server.POST("/login", login)
}