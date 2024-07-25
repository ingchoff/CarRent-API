package routes

import (
	"net/http"

	"example.com/car-rental/models"
	"github.com/gin-gonic/gin"
)

func createService(context *gin.Context) {
  var services map[string]models.Service
  var listServices []models.Service
  err := context.ShouldBindJSON(&services)
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
    context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to add new service."})
    return
  }
  for _, v := range services {
    v.UserID = userId
    listServices = append(listServices, v)
  }
  err = models.Save(listServices)
  if err != nil {
    context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save data."})
    return
  }
  context.JSON(http.StatusOK, gin.H{"message": "Saved data."})
}