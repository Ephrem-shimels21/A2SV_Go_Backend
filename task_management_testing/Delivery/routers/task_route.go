package routers

import (
	"github.com/Ephrem-shimels/A2SV_Go_Backend/task_management_with_cleanArchitecture/Delivery/controllers"
	domain "github.com/Ephrem-shimels/A2SV_Go_Backend/task_management_with_cleanArchitecture/Domain"
	infrastructure "github.com/Ephrem-shimels/A2SV_Go_Backend/task_management_with_cleanArchitecture/Infrastructure"
	repository "github.com/Ephrem-shimels/A2SV_Go_Backend/task_management_with_cleanArchitecture/Repository"
	usecase "github.com/Ephrem-shimels/A2SV_Go_Backend/task_management_with_cleanArchitecture/Usecase"
	"github.com/gin-gonic/gin"
)

func NewTaskRouter(db infrastructure.Database, gin *gin.Engine) {
	tr := repository.NewTaskRepository(db, domain.CollectionTasks)
	tc := &controllers.TaskController{
		TaskUsecase: usecase.NewTaskUsecase(
			tr,
		),
	}
	protectedRoute := gin.Group("")
	publicRoute := gin.Group("")
	protectedRoute.Use(infrastructure.AdminOnlyMiddleware(), infrastructure.JWTAuthMiddleware())
	publicRoute.GET("/tasks", tc.GetTasks)
	publicRoute.GET("/tasks/:id", tc.GetTask)
	protectedRoute.POST("/tasks", tc.CreateTask)
	protectedRoute.PUT("/tasks/:id", tc.UpdateTask)
	protectedRoute.DELETE("/tasks/:id", tc.DeleteTask)

}
