package routers

import (
	"github.com/Ephrem-shimels/A2SV_Go_Backend/task_management_with_cleanArchitecture/Delivery/controllers"
	domain "github.com/Ephrem-shimels/A2SV_Go_Backend/task_management_with_cleanArchitecture/Domain"
	infrastructure "github.com/Ephrem-shimels/A2SV_Go_Backend/task_management_with_cleanArchitecture/Infrastructure"
	repository "github.com/Ephrem-shimels/A2SV_Go_Backend/task_management_with_cleanArchitecture/Repository"
	usecase "github.com/Ephrem-shimels/A2SV_Go_Backend/task_management_with_cleanArchitecture/Usecase"
	"github.com/gin-gonic/gin"
)

func NewTaskRouter(db infrastructure.Database, group *gin.RouterGroup) {
	tr := repository.NewTaskRepository(db, domain.CollectionTasks)
	tc := &controllers.TaskController{
		TaskUsecase: usecase.NewTaskUsecase(
			tr,
		),
	}
	group.GET("/tasks", tc.GetTasks)
	group.GET("/tasks/:id", tc.GetTask)
	group.POST("/tasks", tc.CreateTask)
	group.PUT("/tasks/:id", tc.UpdateTask)
	group.DELETE("/tasks/:id", tc.DeleteTask)

}
