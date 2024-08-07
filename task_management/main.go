package main

import (
	"github.com/Ephrem-shimels/A2SV_Go_Backend/task_management/router"
)

func main() {
	r := router.SetupRouter()
	r.Run(":8080")

}
