package controllers

import (
	"net/http"
	"strconv"

	"github.com/Ephrem-shimels/A2SV_Go_Backend/task_management/data"
	"github.com/Ephrem-shimels/A2SV_Go_Backend/task_management/models"
	"github.com/gin-gonic/gin"
)

func GetTasks(cxt *gin.Context) {
	tasks := data.GetTasks()
	cxt.JSON(http.StatusOK, tasks)
}

func GetTaskByID(cxt *gin.Context) {
	id, err := strconv.Atoi(cxt.Param("id"))
	if err != nil {
		cxt.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}
	task, err := data.GetTask(id)
	if err != nil {
		cxt.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	cxt.JSON(http.StatusOK, task)

}

func CreateTask(cxt *gin.Context) {
	var newTask models.Task

	if err := cxt.ShouldBindJSON(&newTask); err != nil {
		cxt.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdTask := data.CreateTask(newTask)
	cxt.JSON(http.StatusCreated, createdTask)

}

func UpdateTask(cxt *gin.Context) {
	id, err := strconv.Atoi(cxt.Param("id"))

	if err != nil {
		cxt.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	var updatedTask models.Task
	if err := cxt.ShouldBindJSON(&updatedTask); err != nil {
		cxt.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := data.UpdateTask(id, updatedTask)

	if err != nil {
		cxt.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	cxt.JSON(http.StatusOK, task)

}

func DeleteTask(cxt *gin.Context) {
	id, err := strconv.Atoi(cxt.Param("id"))
	if err != nil {
		cxt.JSON(http.StatusBadRequest, gin.H{"error": "invalid task ID"})
		return
	}

	if err := data.DeleteTask(id); err != nil {
		cxt.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	cxt.JSON(http.StatusNoContent, nil)
}
