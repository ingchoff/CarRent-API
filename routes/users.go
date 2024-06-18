package routes

import (
	"net/http"

	"example.com/car-rental/models"
	"example.com/car-rental/utils"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}
	user.ID, err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user."})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Signup successfully"})
}

func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}
	user, err = user.ValidateCredentials()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Username or password is wrong."})
		return
	}
	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	err = models.AddRefreshTokenToWhitelist(user, token)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Login successfully", "token": token})
}

func getUser(context *gin.Context) {
	userId := context.GetUint("UserId")
	user, err := models.GetUserById(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse request user data."})
		return
	}
	context.JSON(http.StatusOK, gin.H{"user": user})
}

func getUserByToken(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	user, err := models.GetUserByToken(token)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse request user data."})
		return
	}
	context.JSON(http.StatusOK, gin.H{"user": user})
}

func revokeToken(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	err := models.RevorkToken(token)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not revorked token."})
		return
	}
	context.JSON(http.StatusOK, gin.H{"msg": "revorked token"})
}