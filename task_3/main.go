package main

import (
	"github.com/Ephrem-shimels21/controllers"
	"github.com/Ephrem-shimels21/models"
	"github.com/Ephrem-shimels21/services"
)

func main() {
	library := services.NewLibrary()

	// Add some sample members
	library.Members[1] = models.Member{ID: 1, Name: "Abebe"}
	library.Members[2] = models.Member{ID: 2, Name: "Kebede"}

	// Run the console application
	controllers.RunConsoleApp(library)
}
