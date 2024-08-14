package usecase

import (
	"context"

	domain "github.com/Ephrem-shimels/A2SV_Go_Backend/task_management_with_cleanArchitecture/Domain"
)

type TaskUsecase struct {
	taskRepository domain.TaskRepository
}

func NewTaskUsecase(taskRepository domain.TaskRepository) domain.TaskUsecase {
	return &TaskUsecase{
		taskRepository: taskRepository,
	}
}

func (taskUc *TaskUsecase) CreateTask(cxt context.Context, task *domain.Task) (domain.Task, error) {
	return taskUc.taskRepository.CreateTask(cxt, task)
}

func (taskUc *TaskUsecase) GetTasks(cxt context.Context) ([]domain.Task, error) {
	return taskUc.taskRepository.GetTasks(cxt)
}

func (taskUc *TaskUsecase) GetTask(cxt context.Context, id string) (domain.Task, error) {
	return taskUc.taskRepository.GetTask(cxt, id)
}

func (taskUc *TaskUsecase) UpdateTask(cxt context.Context, id string, updatedTask *domain.Task) (domain.Task, error) {
	return taskUc.taskRepository.UpdateTask(cxt, id, updatedTask)
}

func (taskUc *TaskUsecase) DeleteTask(cxt context.Context, id string) error {
	return taskUc.taskRepository.DeleteTask(cxt, id)
}
