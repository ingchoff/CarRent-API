package routes

import (
	"example.com/car-rental/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate, middlewares.CORSMiddleware())
	authenticated.GET("/user", getUser)
	authenticated.GET("/user/token", getUserByToken)
	authenticated.POST("/car/new", createCar)
	authenticated.GET("/cars", getCars)
	authenticated.GET("/cars/models-data", getListModelsCar)
	authenticated.GET("/cars/:id", getCar)
	authenticated.GET("/cars/search", searchCar)
	authenticated.PUT("/car/:id", updateCar)
	authenticated.DELETE("/car/:id", deleteCar)
	authenticated.POST("/rental/new", createRental)
	authenticated.GET("/rentals/:cid", getRentals)
	authenticated.GET("/rental/:id", getRental)
	authenticated.PUT("/rental/:id", updateRental)
	authenticated.DELETE("/rental/:id", deleteRental)
	authenticated.GET("/rental/search", searchRental)
	authenticated.POST("/inspection/new", createInspection)
	authenticated.GET("/inspections/:cid", getInspections)
	authenticated.GET("/inspection/search", getInspectionsByType)
	authenticated.GET("/inspection/summary", summaryInspections)
	authenticated.PUT("/inspection/:id", updateInspection)
	authenticated.DELETE("/inspection/:id", deleteInspection)
	authenticated.POST("/service/new", createService)
	authenticated.GET("/services/:cid", getServices)
	authenticated.PUT("/services/edit", updateServices)

	auth := server.Group("/auth")
	auth.Use(middlewares.CORSMiddleware())
	auth.POST("/signup", signup)
	auth.POST("/login", login)
	auth.POST("/revorktoken", revokeToken)
}