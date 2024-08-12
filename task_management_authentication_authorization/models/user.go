package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Username string             `json:"username" bson:"username"`
	Password string             `json:"password" bson:"password"` // This will be hashed
	Role     string             `json:"role" bson:"role"`         // e.g., "admin", "user"
}

type UserDTO struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type PromoteDTO struct {
	UserName string `json:"username"`
}
