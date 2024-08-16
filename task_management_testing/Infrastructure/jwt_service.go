package infrastructure

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"time"

	domain "github.com/Ephrem-shimels/A2SV_Go_Backend/task_management_with_cleanArchitecture/Domain"
	"github.com/golang-jwt/jwt"
)

var Config struct {
	JWTSecretKey string `json:"jwt_secret_key"`
}

func LoadConfig() error {
	file, err := os.Open("../config.json")
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

func GenerateJWT(user *domain.User) (string, error) {
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

	tokenString, err := token.SignedString([]byte(secretKey))

	if err != nil {
		return "", err
	}

	return tokenString, nil

}
