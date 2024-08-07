package data

import (
	"errors"

	"github.com/Ephrem-shimels/A2SV_Go_Backend/task_management/models"
	// "github.com/gin-gonic/gin"
)

var tasks = make(map[int]models.Task)
var currentID = 1

func GetTasks() []models.Task {
	taskList := []models.Task{}

	for _, task := range tasks {
		taskList = append(taskList, task)
	}

	return taskList

}

func GetTask(id int) (models.Task, error) {
	task, ok := tasks[id]

	if !ok {
		return models.Task{}, errors.New("Task not found")
	}

	return task, nil

}

func CreateTask(task models.Task) models.Task {
	task.ID = currentID
	currentID++
	tasks[task.ID] = task
	return task
}

func UpdateTask(id int, updatedTask models.Task) (models.Task, error) {
	_, exist := tasks[id]

	if !exist {
		return models.Task{}, errors.New("Task not found")
	}
	updatedTask.ID = id
	tasks[id] = updatedTask
	return updatedTask, nil

}

func DeleteTask(id int) error {
	_, exist := tasks[id]

	if !exist {
		return errors.New("Task not found")
	}

	delete(tasks, id)
	return nil
}
