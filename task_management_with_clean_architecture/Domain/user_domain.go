package domain

import (
	"context"

	"github.com/Ephrem-shimels/A2SV_Go_Backend/task_management_with_cleanArchitecture/Domain/dtos"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionUsers = "users"
)

type User struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Username string             `json:"username" bson:"username"`
	Password string             `json:"password" bson:"password"` // This will be hashed
	Role     string             `json:"role" bson:"role"`         // e.g., "admin", "user"
}

type UserRepository interface {
	RegisterUser(c context.Context, registerDto dtos.RegisterDto) (*User, error)
	PromoteUser(c context.Context, promoteDto dtos.PromoteDto)
	Login(c context.Context, registerDto dtos.RegisterDto)
}

type UserUsecase interface {
	RegisterUser(c context.Context, userName string, password string) (*User, error)
	PromoteUser(c context.Context, userName string)
	Login(c context.Context, userName string, password string)
}
