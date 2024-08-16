package usecase

import (
	"context"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"

	domain "github.com/Ephrem-shimels/A2SV_Go_Backend/task_management_with_cleanArchitecture/Domain"
	"github.com/Ephrem-shimels/A2SV_Go_Backend/task_management_with_cleanArchitecture/Domain/dtos"
	"github.com/Ephrem-shimels/A2SV_Go_Backend/task_management_with_cleanArchitecture/Repository/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegisterUser(t *testing.T) {
	// Step 1: Create a mock instance of UserRepository
	mockRepo := new(mocks.UserRepository)

	// Step 2: Set up test data
	registerDto := dtos.RegisterDto{
		Username: "testuser",
		Password: "password123",
	}

	// Step 3: Define the expected result user with a placeholder for hashed password
	objectID, err := primitive.ObjectIDFromHex("507f1f77bcf86cd799439011")
	if err != nil {
		t.Fatalf("Failed to create ObjectID from hex: %v", err)
	}

	expectedUser := &domain.User{
		ID:       objectID,
		Username: "testuser",
		Password: "", // Password should not be checked directly here
	}

	// Step 4: Set expectations on the mock
	mockRepo.On("RegisterUser", mock.Anything, registerDto, mock.MatchedBy(func(hashedPassword string) bool {
		// Optionally, you can add custom logic to check if the hashed password matches the expected format
		return len(hashedPassword) > 0 // This just checks if the hashed password is a non-empty string
	})).Return(expectedUser, nil)

	// Step 5: Create the use case and inject the mock repository
	userUsecase := NewUserUsecase(mockRepo)

	// Step 6: Call the method under test
	result, err := userUsecase.RegisterUser(context.Background(), registerDto)

	// Step 7: Assert the results
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, expectedUser.ID, result.ID)
	assert.Equal(t, expectedUser.Username, result.Username)

	// Step 8: Assert that expectations were met
	mockRepo.AssertExpectations(t)
}

func TestPromoteUser(t *testing.T) {
	// Step 1: Create a mock instance of UserRepository
	mockRepo := new(mocks.UserRepository)

	// Step 2: Set up test data and expected result
	promoteDto := dtos.PromoteDto{
		Username: "testuser",
	}

	// Step 3: Set expectations on the mock
	mockRepo.On("PromoteUser", mock.Anything, promoteDto).Return(nil)

	// Step 4: Create the use case and inject the mock repository
	userUsecase := NewUserUsecase(mockRepo)

	// Step 5: Call the method under test
	err := userUsecase.PromoteUser(context.Background(), promoteDto)

	// Step 6: Assert the results
	assert.NoError(t, err)

	// Step 7: Assert that expectations were met
	mockRepo.AssertExpectations(t)
}
