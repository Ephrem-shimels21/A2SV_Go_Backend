package main

import (
	"log"

	"github.com/Ephrem-shimels/A2SV_Go_Backend/task_management_with_cleanArchitecture/Delivery/routers"
	infrastructure "github.com/Ephrem-shimels/A2SV_Go_Backend/task_management_with_cleanArchitecture/Infrastructure"
	"github.com/gin-gonic/gin"
)

func main() {
	gin := gin.Default()

	db, err := infrastructure.ConnectDb()

	if err != nil {
		log.Fatalf("Failed to connect db")
	}

	routers.SetupRouter(db, gin)
	gin.Run(":8080")

}
