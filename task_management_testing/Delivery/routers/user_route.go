package routers

import (
	"github.com/Ephrem-shimels/A2SV_Go_Backend/task_management_with_cleanArchitecture/Delivery/controllers"
	domain "github.com/Ephrem-shimels/A2SV_Go_Backend/task_management_with_cleanArchitecture/Domain"
	infrastructure "github.com/Ephrem-shimels/A2SV_Go_Backend/task_management_with_cleanArchitecture/Infrastructure"
	repository "github.com/Ephrem-shimels/A2SV_Go_Backend/task_management_with_cleanArchitecture/Repository"
	usecase "github.com/Ephrem-shimels/A2SV_Go_Backend/task_management_with_cleanArchitecture/Usecase"
	"github.com/gin-gonic/gin"
)

func NewUserRouter(db infrastructure.Database, gin *gin.Engine) {
	ur := repository.NewUserRepository(db, domain.CollectionUsers)
	uc := &controllers.UserController{
		UserUsecase: usecase.NewUserUsecase(ur),
	}
	groupPublic := gin.Group("")

	groupPublic.POST("/register", uc.RegisterUser)
	groupPublic.POST("/login", uc.Login)

	protectedRoute := gin.Group("")
	protectedRoute.Use(infrastructure.AdminOnlyMiddleware())
	protectedRoute.POST("/promote", uc.PromoteUser)
}
