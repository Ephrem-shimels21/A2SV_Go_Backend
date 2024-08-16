package domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionTasks = "tasks"
)

type Task struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description"`
	DueDate     string             `json:"due_date" bson:"due_date"`
	Status      string             `json:"status" bson:"status"`
}

type TaskRepository interface {
	CreateTask(c context.Context, task *Task) (Task, error)
	GetTasks(c context.Context) ([]Task, error)
	GetTask(c context.Context, id string) (Task, error)
	UpdateTask(c context.Context, id string, updatedTask *Task) (Task, error)
	DeleteTask(c context.Context, id string) error
}

type TaskUsecase interface {
	CreateTask(c context.Context, task *Task) (Task, error)
	GetTasks(c context.Context) ([]Task, error)
	GetTask(c context.Context, id string) (Task, error)
	UpdateTask(c context.Context, id string, updatedTask *Task) (Task, error)
	DeleteTask(c context.Context, id string) error
}
