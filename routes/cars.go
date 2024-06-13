package routes

import (
	"net/http"
	"strconv"

	"example.com/car-rental/models"
	"github.com/gin-gonic/gin"
)

func createCar(context *gin.Context) {
	var car models.Car
	err := context.ShouldBindJSON(&car)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	userId := context.GetUint("UserId")
	user, err := models.GetUserById(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse request user data."})
		return
	}
	if user.Role != "admin" {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to add new car."})
		return
	}
	car.UserID = userId
	err = car.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save data."})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Saved data."})
}

func getCars(context *gin.Context) {
	userId := context.GetUint("UserId")
	cars, err := models.FindAllCars(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch request cars data."})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": cars})
}

func getCar(context *gin.Context) {
	carId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	userId := context.GetUint("UserId")
	if err!= nil {
		context.JSON(http.StatusBadRequest, gin.H{"msg": "Could not parse car id. Try again later."})
		return
	}
	car, err := models.FindCarById(uint(carId))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch request car data."})
		return
	}
	if userId != car.UserID {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to get car data."})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": car})
}

func updateCar(context *gin.Context) {
	carId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err!= nil {
		context.JSON(http.StatusBadRequest, gin.H{"msg": "Could not parse car id. Try again later."})
		return
	}
	userId := context.GetUint("UserId")
	var updateCar models.Car
	err = context.ShouldBindJSON(&updateCar)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"msg": "Could not parse request data."})
		return
	}
	car, err := models.FindCarById(uint(carId))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch request car data."})
		return
	}
	if userId != car.UserID {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to update car."})
		return
	}
	updateCar.UserID = userId
	updateCar.ID = uint(carId)
	err = updateCar.UpdateCar()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"msg": "Could not update the car."})
	}
	context.JSON(http.StatusOK, gin.H{"msg": "car updated successfully."})
}

func deleteCar(context *gin.Context) {
	carId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err!= nil {
		context.JSON(http.StatusBadRequest, gin.H{"msg": "Could not parse car id. Try again later."})
		return
	}
	userId := context.GetUint("UserId")
	car, err := models.FindCarById(uint(carId))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch request car data."})
		return
	}
	if userId != car.UserID {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to delete car."})
		return
	}
	err = models.DeleteCarById(uint(carId))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete data."})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Deleted data."})
}