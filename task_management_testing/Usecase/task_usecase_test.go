package usecase

import (
	"context"
	"testing"

	domain "github.com/Ephrem-shimels/A2SV_Go_Backend/task_management_with_cleanArchitecture/Domain"
	"github.com/Ephrem-shimels/A2SV_Go_Backend/task_management_with_cleanArchitecture/Repository/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCreateTask(t *testing.T) {
	mockRepo := mocks.NewTaskRepository(t)
	taskUc := NewTaskUsecase(mockRepo)
	taskID, err := primitive.ObjectIDFromHex("507f1f77bcf86cd799439011")

	if err != nil {
		t.Fatalf("Failed to create ObjectID from hex: %v", err)
	}

	task := &domain.Task{ID: taskID, Title: "Test Task"}
	mockRepo.On("CreateTask", mock.Anything, task).Return(*task, nil)

	result, err := taskUc.CreateTask(context.Background(), task)

	assert.NoError(t, err)
	assert.Equal(t, *task, result)
	mockRepo.AssertExpectations(t)
}

func TestGetTasks(t *testing.T) {
	mockRepo := mocks.NewTaskRepository(t)
	taskUc := NewTaskUsecase(mockRepo)
	taskID, err := primitive.ObjectIDFromHex("507f1f77bcf86cd799439011")
	if err != nil {
		t.Fatalf("Failed to create ObjectID from hex: %v", err)
	}
	taskID2, err := primitive.ObjectIDFromHex("507f1f77bcf86cd799439012")

	if err != nil {
		t.Fatalf("Failed to create ObjectID from hex: %v", err)
	}

	tasks := []domain.Task{
		{ID: taskID, Title: "Task 1"},
		{ID: taskID2, Title: "Task 2"},
	}
	mockRepo.On("GetTasks", mock.Anything).Return(tasks, nil)

	result, err := taskUc.GetTasks(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, tasks, result)
	mockRepo.AssertExpectations(t)
}

func TestGetTask(t *testing.T) {
	mockRepo := mocks.NewTaskRepository(t)
	taskUc := NewTaskUsecase(mockRepo)

	taskID, err := primitive.ObjectIDFromHex("507f1f77bcf86cd799439011")
	if err != nil {
		t.Fatalf("Failed to create ObjectID from hex: %v", err)
	}

	task := domain.Task{ID: taskID, Title: "Test Task"}
	mockRepo.On("GetTask", mock.Anything, "507f1f77bcf86cd799439011").Return(task, nil)

	result, err := taskUc.GetTask(context.Background(), "507f1f77bcf86cd799439011")

	assert.NoError(t, err)
	assert.Equal(t, task, result)
	mockRepo.AssertExpectations(t)
}

func TestUpdateTask(t *testing.T) {
	mockRepo := mocks.NewTaskRepository(t)
	taskUc := NewTaskUsecase(mockRepo)

	taskID, err := primitive.ObjectIDFromHex("507f1f77bcf86cd799439011")

	if err != nil {
		t.Fatalf("Failed to create ObjectID from hex: %v", err)
	}

	task := &domain.Task{ID: taskID, Title: "Test Task"}
	mockRepo.On("UpdateTask", mock.Anything, "1", task).Return(*task, nil)

	result, err := taskUc.UpdateTask(context.Background(), "1", task)

	assert.NoError(t, err)
	assert.Equal(t, *task, result)
	mockRepo.AssertExpectations(t)
}

func TestDeleteTask(t *testing.T) {
	mockRepo := mocks.NewTaskRepository(t)
	taskUc := NewTaskUsecase(mockRepo)

	mockRepo.On("DeleteTask", mock.Anything, "1").Return(nil)

	err := taskUc.DeleteTask(context.Background(), "1")

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
