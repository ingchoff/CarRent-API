package routes

import (
	"net/http"
	"strconv"

	"example.com/car-rental/models"
	"github.com/gin-gonic/gin"
)

func createInspection(context *gin.Context) {
	var ins models.Inspection
	err := context.ShouldBindJSON(&ins)
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
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to add new inspection."})
		return
	}
	ins.UserID = userId
	err = ins.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save data."})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Saved data."})
}

func getInspections(context *gin.Context) {
	carId, err := strconv.ParseInt(context.Param("cid"), 10, 64)
	if err!= nil {
		context.JSON(http.StatusBadRequest, gin.H{"msg": "Could not parse car id. Try again later."})
		return
	}
	inspections, err := models.FindAllInspections(uint(carId))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch request inspections data."})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": inspections})
}

func getInspectionsByType(context *gin.Context) {
	// service := context.DefaultQuery("service", "")
	services := context.Request.URL.Query()
	if (!services.Has("carid") || !services.Has("service")) {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Query Params Required!."})
		return
	} else {
		inspections, err := models.FindInsByType(services.Get("service"), services.Get("carid"))
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch request inspections data."})
			return
		}
		context.JSON(http.StatusOK, gin.H{"data": inspections})
	}
}

func summaryInspections(context *gin.Context) {
	service := context.Request.URL.Query()
	if (!service.Has("carid")) {
		context.JSON(http.StatusBadRequest, gin.H{"message": "query params carid required!."})
	} else {
		sum, err := models.LatestInsByCar(service.Get("carid"))
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch request inspections data."})
			return
		}
		context.JSON(http.StatusOK, gin.H{"data": sum})
	}
}

func updateInspection(context *gin.Context) {
	inspectionId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err!= nil {
		context.JSON(http.StatusBadRequest, gin.H{"msg": "Could not parse inspection id. Try again later."})
		return
	}
	userId := context.GetUint("UserId")
	var updateIns models.Inspection
	err = context.ShouldBindJSON(&updateIns)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"msg": "Could not parse request data."})
		return
	}
	car, err := models.FindCarById(updateIns.CarID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch request car data."})
		return
	}
	if userId != car.UserID {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to update inspection."})
		return
	}
	updateIns.ID = uint(inspectionId)
	err = updateIns.UpdateIns()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"msg": "Could not update the inspection."})
	}
	context.JSON(http.StatusOK, gin.H{"msg": "car updated successfully."})
}

func deleteInspection(context *gin.Context) {
	inspectionId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err!= nil {
		context.JSON(http.StatusBadRequest, gin.H{"msg": "Could not parse car id. Try again later."})
		return
	}
	userId := context.GetUint("UserId")
	var updateIns models.Inspection
	err = context.ShouldBindJSON(&updateIns)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"msg": "Could not parse request data."})
		return
	}
	car, err := models.FindCarById(updateIns.CarID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch request car data."})
		return
	}
	if userId != car.UserID {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to delete inspection."})
		return
	}
	err = models.DeleteInsById(uint(inspectionId))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete data."})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Deleted data."})
}