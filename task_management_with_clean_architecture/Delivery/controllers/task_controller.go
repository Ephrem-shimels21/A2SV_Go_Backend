package controllers

import (
	"net/http"

	domain "github.com/Ephrem-shimels/A2SV_Go_Backend/task_management_with_cleanArchitecture/Domain"
	"github.com/gin-gonic/gin"
)

type TaskController struct {
	TaskUsecase domain.TaskUsecase
}

func (taskcont *TaskController) CreateTask(cxt *gin.Context) {
	var newTask domain.Task

	err := cxt.ShouldBindJSON(&newTask)

	if err != nil {
		cxt.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdTask, err := taskcont.TaskUsecase.CreateTask(cxt, &newTask)

	if err != nil {
		cxt.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	cxt.JSON(http.StatusOK, createdTask)

}

func (taskCont *TaskController) GetTasks(cxt *gin.Context) {
	tasks, err := taskCont.TaskUsecase.GetTasks(cxt)

	if err != nil {
		cxt.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

	}

	cxt.JSON(http.StatusOK, tasks)

}

func (taskCont *TaskController) GetTask(cxt *gin.Context) {
	taskID := cxt.Param("id")

	task, err := taskCont.TaskUsecase.GetTask(cxt, taskID)

	if err != nil {
		cxt.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	cxt.JSON(http.StatusOK, task)

}

func (taskCont *TaskController) UpdateTask(cxt *gin.Context) {
	taskID := cxt.GetString("id")

	var updatedTask domain.Task

	if err := cxt.ShouldBindJSON(&updatedTask); err != nil {
		cxt.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := taskCont.TaskUsecase.UpdateTask(cxt, taskID, &updatedTask)

	if err != nil {
		cxt.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"error": err.Error()}})
	}

	cxt.JSON(http.StatusOK, task)

}

func (taskCont *TaskController) DeleteTask(cxt *gin.Context) {
	taskID := cxt.GetString("id")

	err := taskCont.TaskUsecase.DeleteTask(cxt, taskID)

	if err != nil {
		cxt.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	cxt.JSON(http.StatusOK, nil)
}
