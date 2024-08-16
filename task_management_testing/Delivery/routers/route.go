package routers

import (
	infrastructure "github.com/Ephrem-shimels/A2SV_Go_Backend/task_management_with_cleanArchitecture/Infrastructure"
	"github.com/gin-gonic/gin"
)

func SetupRouter(db infrastructure.Database, gin *gin.Engine) {
	NewUserRouter(db, gin)
	NewTaskRouter(db, gin)

	// NewTaskRouter(db, router)

}
