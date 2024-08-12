package controllers

import (
	"net/http"

	"github.com/Ephrem-shimels/A2SV_Go_Backend/enhancing_task_management/data"
	"github.com/Ephrem-shimels/A2SV_Go_Backend/enhancing_task_management/models"
	"github.com/gin-gonic/gin"
)

func GetTasks(cxt *gin.Context) {
	tasks, err := data.GetTasks()
	if err != nil {
		cxt.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	cxt.JSON(http.StatusOK, tasks)
}

func GetTaskByID(cxt *gin.Context) {
	id := cxt.Param("id")

	task, err := data.GetTask(id)
	if err != nil {
		cxt.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
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

	createdTask, err := data.CreateTask(newTask)
	if err != nil {
		cxt.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return

	}
	cxt.JSON(http.StatusCreated, createdTask)

}

func UpdateTask(cxt *gin.Context) {
	id := cxt.Param("id")

	var updatedTask models.Task
	if err := cxt.ShouldBindJSON(&updatedTask); err != nil {
		cxt.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := data.UpdateTask(id, updatedTask)

	if err != nil {
		cxt.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update the task"})
		return
	}

	cxt.JSON(http.StatusOK, task)

}

func DeleteTask(cxt *gin.Context) {
	id := cxt.Param("id")

	if err := data.DeleteTask(id); err != nil {
		cxt.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	cxt.JSON(http.StatusNoContent, nil)
}
