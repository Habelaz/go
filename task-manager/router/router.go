package router

import (
	"task-manager/controllers"
	"task-manager/middleware"

	"github.com/gin-gonic/gin"
)

func InitializeRouter(){
	router := gin.Default()

	router.POST("/register",controllers.Register)
	router.POST("/login",controllers.Login)

	
	router.GET("/tasks",middleware.AuthMiddleware(),controllers.GetTasks)
	router.GET("/tasks/:id",middleware.AuthMiddleware(),controllers.GetSpecificTask)
	router.PUT("/tasks/:id",middleware.AuthMiddleware(),middleware.AdminOnly(),controllers.UpdateTask)
	router.POST("/tasks",middleware.AuthMiddleware(),controllers.CreateTask)
	router.DELETE("/tasks/:id",middleware.AuthMiddleware(),middleware.AdminOnly(),controllers.DeleteTask)

	router.Run("localhost:8080")
}