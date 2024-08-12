package router

import (
	"github.com/Ephrem-shimels/A2SV_Go_Backend/enhancing_task_management/controllers"
	"github.com/Ephrem-shimels/A2SV_Go_Backend/enhancing_task_management/middleware"
	"github.com/gin-gonic/gin"
)

// func SetupRouter() *gin.Engine {
// 	router := gin.Default()

// 	router.GET("/tasks", controllers.GetTasks)
// 	router.GET("/tasks/:id", controllers.GetTaskByID)
// 	router.POST("/tasks", controllers.CreateTask)
// 	router.PUT("/tasks/:id", controllers.UpdateTask)
// 	router.DELETE("/tasks/:id", controllers.DeleteTask)

// 	return router
// }

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Public routes
	r.POST("/register", controllers.RegisterUser)
	r.POST("/login", controllers.LoginUser)

	// Protected routes
	authorized := r.Group("/")
	authorized.Use(middleware.JWTAuthMiddleware())
	{
		authorized.GET("/tasks", controllers.GetTasks)
		authorized.GET("/tasks/:id", controllers.GetTaskByID)
		authorized.Use(middleware.AdminOnlyMiddleware())
		{
			authorized.POST("/tasks", controllers.CreateTask)
			authorized.PUT("/tasks/:id", controllers.UpdateTask)
			authorized.DELETE("/tasks/:id", controllers.DeleteTask)
			authorized.POST("/promote", controllers.PromoteUser)
		}
	}

	return r
}
