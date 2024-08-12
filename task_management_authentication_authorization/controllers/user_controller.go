package controllers

import (
	"net/http"

	"github.com/Ephrem-shimels/A2SV_Go_Backend/enhancing_task_management/data"
	"github.com/Ephrem-shimels/A2SV_Go_Backend/enhancing_task_management/models"
	"github.com/gin-gonic/gin"
)

func RegisterUser(cxt *gin.Context) {
	var userDTO models.UserDTO

	if err := cxt.ShouldBindJSON(&userDTO); err != nil {
		cxt.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := data.RegisterUser(userDTO)

	if err != nil {
		cxt.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	cxt.JSON(http.StatusOK, user)

}

func LoginUser(cxt *gin.Context) {
	var userDTO models.UserDTO

	if err := cxt.ShouldBindJSON(&userDTO); err != nil {
		cxt.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := data.AuthenticateUser(userDTO)

	if err != nil {
		cxt.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	tokenString, err := data.GenerateJWT(user)

	if err != nil {
		cxt.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}

	cxt.JSON(http.StatusOK, gin.H{"user": user, "token": tokenString})
}

func PromoteUser(cxt *gin.Context) {
	var promoteDTO models.PromoteDTO

	if err := cxt.ShouldBindJSON(&promoteDTO); err != nil {
		cxt.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := data.PromoteUser(promoteDTO)

	if err != nil {
		cxt.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cxt.JSON(http.StatusOK, gin.H{"message": "User promoted to admin, successfully"})
}
