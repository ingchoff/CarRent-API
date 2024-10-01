package routes

import (
	"net/http"
	"strconv"

	"example.com/car-rental/models"
	"github.com/gin-gonic/gin"
)

func createRental(context *gin.Context) {
	var rental models.Rental
	err := context.ShouldBindJSON(&rental)
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
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to add new rental."})
		return
	}
	rental.UserID = userId
	err = rental.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save data."})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Saved data."})
}

func getRentals(context *gin.Context) {
	userId := context.GetUint("UserId")
	cid, err := strconv.ParseInt(context.Param("cid"), 10, 64)
	if err!= nil {
		context.JSON(http.StatusBadRequest, gin.H{"msg": "Could not parse event id. Try again later."})
		return
	}
	rentals, err := models.FindAllRentals(userId, uint(cid))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch request rentals data."})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": rentals})
}

func getRental(context *gin.Context) {
	rentalId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	userId := context.GetUint("UserId")
	if err!= nil {
		context.JSON(http.StatusBadRequest, gin.H{"msg": "Could not parse event id. Try again later."})
		return
	}
	rental, err := models.FindRentalById(uint(rentalId))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch request rental data."})
		return
	}
	if userId != rental.UserID {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to get rental data."})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": rental})
}

func updateRental(context *gin.Context) {
	rentalId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err!= nil {
		context.JSON(http.StatusBadRequest, gin.H{"msg": "Could not parse event id. Try again later."})
		return
	}
	userId := context.GetUint("UserId")
	var updateRental models.Rental
	err = context.ShouldBindJSON(&updateRental)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"msg": "Could not parse request data."})
		return
	}
	rental, err := models.FindRentalById(uint(rentalId))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch request rental data."})
		return
	}
	if userId != rental.UserID {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to update rental."})
		return
	}
	updateRental.UserID = userId
	updateRental.ID = uint(rentalId)
	err = updateRental.UpdateRental()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"msg": "Could not update the rental."})
		return
	}
	context.JSON(http.StatusOK, gin.H{"msg": "Event updated successfully."})
}

func deleteRental(context *gin.Context) {
	rentalId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err!= nil {
		context.JSON(http.StatusBadRequest, gin.H{"msg": "Could not parse rental id. Try again later."})
		return
	}
	userId := context.GetUint("UserId")
	rental, err := models.FindRentalById(uint(rentalId))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch request rental data."})
		return
	}
	if userId != rental.UserID {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to delete rental."})
		return
	}
	err = models.DeleteRentalById(uint(rentalId))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete data."})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Deleted data."})
}

func searchRental(context *gin.Context) {
	params := context.Request.URL.Query()
	userId := context.GetUint("UserId")
	if (params.Has("cid") && params.Has("name")) {
		rentals, err := models.FindRentalByCondition(userId, params.Get("name"), params.Get("nid"), params.Get("start"), params.Get("end"), params.Get("cid"))
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get rentals data."})
			return
		}
		context.JSON(http.StatusOK, gin.H{"data": rentals})
	} else {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Query Params Required!."})
		return
	}
}