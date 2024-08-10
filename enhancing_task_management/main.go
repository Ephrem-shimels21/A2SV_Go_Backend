package main

import (
	"github.com/Ephrem-shimels/A2SV_Go_Backend/enhancing_task_management/data"
	"github.com/Ephrem-shimels/A2SV_Go_Backend/enhancing_task_management/router"
)

func main() {
	data.ConnectDb()
	r := router.SetupRouter()
	r.Run(":8080")

}
