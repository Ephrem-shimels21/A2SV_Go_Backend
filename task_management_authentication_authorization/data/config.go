package data

// import (
// 	"encoding/json"
// 	"os"
// )

// type Config struct {
// 	JWTSecretKey string `json:"jwt_secret_key"`
// }

// func loadConfig(config *Config) error {
// 	file, err := os.Open("config.json")
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()

// 	decoder := json.NewDecoder(file)
// 	err = decoder.Decode(config)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
