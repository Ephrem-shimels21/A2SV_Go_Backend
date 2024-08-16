package repository

import (
	"context"
	"testing"

	domain "github.com/Ephrem-shimels/A2SV_Go_Backend/task_management_with_cleanArchitecture/Domain"
	"github.com/Ephrem-shimels/A2SV_Go_Backend/task_management_with_cleanArchitecture/Infrastructure/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestCreateTask(t *testing.T) {
	// Mock setup
	mockDb := new(mocks.Database)
	mockCollection := new(mocks.Collection)

	// Test data
	task := &domain.Task{
		Title:       "Test Task",
		Description: "Test Description",
		DueDate:     "2024-12-31",
		Status:      "Pending",
	}

	// Mocking the collection and database
	mockDb.On("Collection", "tasks").Return(mockCollection)
	mockCollection.On("InsertOne", mock.Anything, task).Return(&mongo.InsertOneResult{}, nil)

	// Create repository
	repo := NewTaskRepository(mockDb, "tasks")

	// Call method
	result, err := repo.CreateTask(context.Background(), task)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, *task, result)

	// Verify expectations
	mockDb.AssertExpectations(t)
	mockCollection.AssertExpectations(t)
}

func TestGetTasks(t *testing.T) {
	// Mock setup
	mockDb := new(mocks.Database)
	mockCollection := new(mocks.Collection)
	mockCursor := new(mocks.Cursor)

	// Test data
	tasks := []domain.Task{
		{ID: primitive.NewObjectID(), Title: "Task 1", Description: "Description 1", DueDate: "2024-12-31", Status: "Pending"},
		{ID: primitive.NewObjectID(), Title: "Task 2", Description: "Description 2", DueDate: "2024-11-30", Status: "Completed"},
	}

	// Mocking the collection and cursor
	mockDb.On("Collection", "tasks").Return(mockCollection)
	mockCollection.On("Find", context.Background(), bson.D{}).Return(mockCursor, nil)
	mockCursor.On("All", context.Background(), mock.Anything).Run(func(args mock.Arguments) {
		ptr := args.Get(1).(*[]domain.Task)
		*ptr = tasks
	}).Return(nil)

	// Create repository
	repo := NewTaskRepository(mockDb, "tasks")

	// Call method
	result, err := repo.GetTasks(context.Background())

	// Assertions
	assert.NoError(t, err)
	assert.ElementsMatch(t, tasks, result)

	// Verify expectations
	mockDb.AssertExpectations(t)
	mockCollection.AssertExpectations(t)
	mockCursor.AssertExpectations(t)
}

func TestGetTask(t *testing.T) {
	// Create mock objects
	mockDb := new(mocks.Database)
	mockCollection := new(mocks.Collection)
	mockSingleResult := mocks.NewSingleResult(t)

	// Test data
	taskID := primitive.NewObjectID()
	task := domain.Task{
		ID:          taskID,
		Title:       "Task 1",
		Description: "Description 1",
		DueDate:     "2024-12-31",
		Status:      "Pending",
	}

	// Setup expectations
	mockDb.On("Collection", "tasks").Return(mockCollection)
	mockCollection.On("FindOne", context.Background(), bson.M{"_id": taskID}).Return(mockSingleResult)
	mockSingleResult.On("Decode", mock.Anything).Run(func(args mock.Arguments) {
		ptr := args.Get(0).(*domain.Task)
		*ptr = task
	}).Return(nil)

	// Create repository
	repo := NewTaskRepository(mockDb, "tasks")

	// Call method
	result, err := repo.GetTask(context.Background(), taskID.Hex())

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, task, result)

	// Verify expectations
	mockDb.AssertExpectations(t)
	mockCollection.AssertExpectations(t)
	mockSingleResult.AssertExpectations(t)
}

func TestUpdateTask(t *testing.T) {
	// Mock setup
	mockDb := new(mocks.Database)
	mockCollection := new(mocks.Collection)

	// Test data
	taskID := primitive.NewObjectID()
	updatedTask := &domain.Task{
		Title:       "Updated Task",
		Description: "Updated Description",
		DueDate:     "2024-12-31",
		Status:      "Completed",
	}

	filter := bson.M{"_id": taskID}
	update := bson.M{
		"$set": bson.M{
			"title":       updatedTask.Title,
			"description": updatedTask.Description,
			"due_date":    updatedTask.DueDate,
			"status":      updatedTask.Status,
		},
	}

	// Mocking the collection and database
	mockDb.On("Collection", "tasks").Return(mockCollection)
	mockCollection.On("UpdateOne", context.Background(), filter, update).Return(&mongo.UpdateResult{}, nil)

	// Create repository
	repo := NewTaskRepository(mockDb, "tasks")

	// Call method
	result, err := repo.UpdateTask(context.Background(), taskID.Hex(), updatedTask)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, *updatedTask, result)

	// Verify expectations
	mockDb.AssertExpectations(t)
	mockCollection.AssertExpectations(t)
}
func TestDeleteTask(t *testing.T) {
	// Create a new mock database and collection
	db := new(mocks.Database)
	collection := new(mocks.Collection)

	// Define the mock return values for the Collection's DeleteOne method
	collection.On("DeleteOne", mock.Anything, mock.MatchedBy(func(arg interface{}) bool {
		// Assert that the argument is a bson.M and has the correct _id field
		m, ok := arg.(bson.M)
		if !ok {
			return false
		}
		_, hasID := m["_id"]
		return hasID
	})).Return(int64(1), nil)

	// Set up the Database mock to return the Collection mock
	db.On("Collection", "tasks").Return(collection)

	// Create an instance of the repository with the mock database
	repo := &taskRepository{
		database:   db,
		collection: "tasks",
	}

	// Define a valid ObjectID for testing
	objID := primitive.NewObjectID()
	id := objID.Hex()

	// Call the DeleteTask method
	err := repo.DeleteTask(context.Background(), id)

	// Assert that there was no error
	assert.NoError(t, err)

	// Assert that the Collection's DeleteOne method was called with the correct arguments
	collection.AssertCalled(t, "DeleteOne", mock.Anything, bson.M{"_id": objID})
}
