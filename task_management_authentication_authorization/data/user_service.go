package data

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"os"
	"time"

	"github.com/Ephrem-shimels/A2SV_Go_Backend/enhancing_task_management/models"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var Config struct {
	JWTSecretKey string `json:"jwt_secret_key"`
}

func LoadConfig() error {
	file, err := os.Open("config.json")
	if err != nil {
		log.Println("Error opening config file:", err)
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	err = decoder.Decode(&Config)
	if err != nil {
		log.Println("Error decoding config file:", err)
		return err
	}

	log.Println("Config loaded successfully:", Config.JWTSecretKey)
	return nil
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(hashedPassword), err
}

func CompareHashPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func AuthenticateUser(userDTO models.UserDTO) (*models.User, error) {
	var existingUser models.User

	err := UserCollection.FindOne(context.TODO(), bson.M{"username": userDTO.UserName}).Decode(&existingUser)

	if err != nil {
		return nil, errors.New("user not found")
	}
	// newHashedPassword, err := HashPassword(userDTO.Password)

	// if err != nil {
	// 	return nil, errors.New(err.Error())
	// }

	err = bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(userDTO.Password))

	if err != nil {
		return nil, errors.New("invalid Credentials")
	}

	return &existingUser, nil

}

func GenerateJWT(user *models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":   user.ID.Hex(),
		"username": user.Username,
		"role":     user.Role,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	})

	secretKey := Config.JWTSecretKey

	if secretKey == "" {
		return "", errors.New("JWT secret key not set")
	}

	tokenSting, err := token.SignedString([]byte(secretKey))

	if err != nil {
		return "", err
	}

	return tokenSting, nil
}

func RegisterUser(userDTO models.UserDTO) (*models.User, error) {
	var existingUser models.User
	err := UserCollection.FindOne(context.TODO(), bson.M{"username": userDTO.UserName}).Decode(&existingUser)

	if err != mongo.ErrNoDocuments {
		return nil, errors.New("username already exists")
	}

	hashedPassword, err := HashPassword(userDTO.Password)

	if err != nil {
		return nil, errors.New("please try again, some internal error")
	}

	role := "user"
	userCount, err := UserCollection.CountDocuments(context.Background(), bson.M{})

	if err == nil && userCount == 0 {
		role = "admin"
	}

	user := models.User{
		Username: userDTO.UserName,
		Password: hashedPassword,
		Role:     role,
	}

	_, err = UserCollection.InsertOne(context.TODO(), user)

	if err != nil {
		return nil, errors.New("fialed to create user")
	}

	return &user, nil

}

func PromoteUser(promoteDTO models.PromoteDTO) error {
	filter := bson.M{"username": promoteDTO.UserName}
	update := bson.M{"$set": bson.M{"role": "admin"}}

	_, err := UserCollection.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		return err
	}

	return nil
}
