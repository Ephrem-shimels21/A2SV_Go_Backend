package repository

import (
	"context"
	"testing"

	domain "github.com/Ephrem-shimels/A2SV_Go_Backend/task_management_with_cleanArchitecture/Domain"
	"github.com/Ephrem-shimels/A2SV_Go_Backend/task_management_with_cleanArchitecture/Domain/dtos"
	"github.com/Ephrem-shimels/A2SV_Go_Backend/task_management_with_cleanArchitecture/Infrastructure/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestFindUser(t *testing.T) {
	// Mock setup
	mockDb := new(mocks.Database)
	mockCollection := new(mocks.Collection)
	mockSingleResult := new(mocks.SingleResult)

	// Define the expected user
	expectedUser := &domain.User{
		Username: "testuser",
		Password: "hashedpassword123",
		Role:     "user",
	}

	// Mocking the SingleResult's Decode method
	mockSingleResult.On("Decode", mock.Anything).Run(func(args mock.Arguments) {
		arg := args.Get(0).(*domain.User)
		*arg = *expectedUser
	}).Return(nil) // Ensure Decode returns nil for no error

	// Mocking the collection and database
	mockDb.On("Collection", "users").Return(mockCollection)
	mockCollection.On("FindOne", mock.Anything, bson.M{"username": "testuser"}).
		Return(mockSingleResult)

	// Create repository
	repo := &UserRepository{
		database:   mockDb,
		collection: "users",
	}

	// Call method
	result, err := repo.FindUser(context.Background(), dtos.RegisterDto{Username: "testuser"})

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, expectedUser.Username, result.Username)
	assert.Equal(t, expectedUser.Password, result.Password)

	// Verify expectations
	mockDb.AssertExpectations(t)
	mockCollection.AssertExpectations(t)
	mockSingleResult.AssertExpectations(t)
}

func TestRegisterUser(t *testing.T) {
	// Mock setup
	mockDb := new(mocks.Database)
	mockCollection := new(mocks.Collection)
	mockSingleResult := new(mocks.SingleResult)

	// Test data
	registerDto := dtos.RegisterDto{
		Username: "newuser",
		Password: "password123",
	}
	hashPassword := "hashedpassword123"

	// Mocking the collection and database
	mockDb.On("Collection", "users").Return(mockCollection)
	mockCollection.On("FindOne", mock.Anything, bson.M{"username": "newuser"}).
		Return(mockSingleResult)

	// Mocking the behavior of FindOne method to return the mockSingleResult
	mockCollection.On("CountDocuments", mock.Anything, bson.M{}).Return(int64(0), nil)
	mockCollection.On("InsertOne", mock.Anything, mock.Anything).Return(&mongo.InsertOneResult{}, nil)

	// Mocking the behavior of Decode method on mockSingleResult
	mockSingleResult.On("Decode", mock.Anything).Run(func(args mock.Arguments) {
		arg := args.Get(0).(*domain.User)
		*arg = domain.User{} // Simulating no user found by leaving it empty
	}).Return(mongo.ErrNoDocuments) // Simulating that no user was found

	// Create repository
	repo := &UserRepository{
		database:   mockDb,
		collection: "users",
	}

	// Call method
	result, err := repo.RegisterUser(context.Background(), registerDto, hashPassword)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, registerDto.Username, result.Username)
	assert.Equal(t, hashPassword, result.Password)

	// Verify expectations
	mockDb.AssertExpectations(t)
	mockCollection.AssertExpectations(t)
	mockSingleResult.AssertExpectations(t)
}

func TestPromoteUser(t *testing.T) {
	// Mock setup
	mockDb := new(mocks.Database)
	mockCollection := new(mocks.Collection)

	// Test data
	promoteDto := dtos.PromoteDto{
		Username: "testuser",
	}

	// Mocking the collection and database
	mockDb.On("Collection", "users").Return(mockCollection)
	mockCollection.On("UpdateOne", mock.Anything, bson.M{"username": "testuser"}, bson.M{"$set": bson.M{"role": "admin"}}).Return(&mongo.UpdateResult{}, nil)

	// Create repository
	repo := &UserRepository{
		database:   mockDb,
		collection: "users",
	}

	// Call method
	err := repo.PromoteUser(context.Background(), promoteDto)

	// Assertions
	assert.NoError(t, err)

	// Verify expectations
	mockDb.AssertExpectations(t)
	mockCollection.AssertExpectations(t)
}
