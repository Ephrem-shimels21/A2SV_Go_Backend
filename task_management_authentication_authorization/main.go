package main

import (
	"log"

	"github.com/Ephrem-shimels/A2SV_Go_Backend/enhancing_task_management/data"
	"github.com/Ephrem-shimels/A2SV_Go_Backend/enhancing_task_management/router"
)

func main() {

	err := data.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	data.ConnectDb()
	r := router.SetupRouter()
	r.Run(":8080")

}
