package repository

import (
	"context"

	domain "github.com/Ephrem-shimels/A2SV_Go_Backend/task_management_with_cleanArchitecture/Domain"
	"github.com/Ephrem-shimels/A2SV_Go_Backend/task_management_with_cleanArchitecture/Domain/dtos"
	infrastructure "github.com/Ephrem-shimels/A2SV_Go_Backend/task_management_with_cleanArchitecture/Infrastructure"
	"go.mongodb.org/mongo-driver/bson"
)

type UserRepository struct {
	database   infrastructure.Database
	collection string
}

func NewUserRepository(db infrastructure.Database, collection string) domain.UserRepository {
	return &UserRepository{
		database:   db,
		collection: collection,
	}
}

func (userRepo *UserRepository) FindUser(c context.Context, registerDto dtos.RegisterDto) (*domain.User, error) {
	var existingUser domain.User
	collection := userRepo.database.Collection(userRepo.collection)

	err := collection.FindOne(c, bson.M{"username": registerDto.Username}).Decode(&existingUser)

	if err != nil {
		return &domain.User{}, err
	}

	return &existingUser, nil

}

func (userRepo *UserRepository) countUsers(c context.Context) (int64, error) {
	collection := userRepo.database.Collection(userRepo.collection)

	userCount, err := collection.CountDocuments(c, bson.M{})

	return userCount, err

}

func (userRepo *UserRepository) RegisterUser(c context.Context, registerDto dtos.RegisterDto, hashPassword string) (*domain.User, error) {
	var existingUser domain.User
	collection := userRepo.database.Collection(userRepo.collection)

	err := collection.FindOne(c, bson.M{"username": registerDto.Username}).Decode(&existingUser)

	if err == nil {
		return &domain.User{}, err
	}

	role := "user"

	userCount, err := collection.CountDocuments(c, bson.M{})

	if err == nil && userCount == 0 {
		role = "admin"
	}

	newUser := domain.User{
		Username: registerDto.Username,
		Password: hashPassword,
		Role:     role,
	}

	_, err = collection.InsertOne(c, newUser)

	if err != nil {
		return &newUser, err
	}

	return &newUser, nil

}

func (userRepo *UserRepository) Login(c context.Context, registerDto dtos.RegisterDto, token string) (string, error) {
	// collection := userRepo.database.Collection(userRepo.collection)

	// var user domain.User

	// err := collection.FindOne(c, bson.M{"username": registerDto.Username}).Decode(&user)

	// if err != nil {
	// 	return "", err
	// }

	// err = infrastructure.CompareHashPassword(user.Password, registerDto.Password)

	// if err != nil {
	// 	return "", err
	// }

	// tokenString, err := infrastructure.GenerateJWT(&user)

	// if err != nil {
	// 	return "", err
	// }

	return token, nil

}

func (userRepo *UserRepository) PromoteUser(c context.Context, promoteDto dtos.PromoteDto) error {
	collection := userRepo.database.Collection(userRepo.collection)

	filter := bson.M{"username": promoteDto.Username}
	update := bson.M{"$set": bson.M{"role": "admin"}}

	_, err := collection.UpdateOne(c, filter, update)

	if err != nil {
		return err
	}

	return nil

}
