package controllers

import (
	"net/http"

	domain "github.com/Ephrem-shimels/A2SV_Go_Backend/task_management_with_cleanArchitecture/Domain"
	"github.com/Ephrem-shimels/A2SV_Go_Backend/task_management_with_cleanArchitecture/Domain/dtos"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserUsecase domain.UserUsecase
}

func (userCont *UserController) RegisterUser(cxt *gin.Context) {
	var newUser dtos.RegisterDto

	err := cxt.ShouldBindJSON(&newUser)

	if err != nil {
		cxt.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdUser, err := userCont.UserUsecase.RegisterUser(cxt, newUser)

	if err != nil {
		cxt.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	cxt.JSON(http.StatusCreated, createdUser)

}

func (userCont *UserController) Login(cxt *gin.Context) {
	var userDto dtos.RegisterDto
	err := cxt.ShouldBindJSON(&userDto)

	if err != nil {
		cxt.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := userCont.UserUsecase.Login(cxt, userDto)

	if err != nil {
		cxt.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	cxt.JSON(http.StatusOK, gin.H{"token": token})

}

func (userCont *UserController) PromoteUser(cxt *gin.Context) {
	var promoteDto dtos.PromoteDto

	err := cxt.ShouldBindJSON(&promoteDto)

	if err != nil {
		cxt.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}

	err = userCont.UserUsecase.PromoteUser(cxt, promoteDto)

	if err != nil {
		cxt.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	cxt.JSON(http.StatusOK, gin.H{"message": "The Role successfuly updated"})

}
