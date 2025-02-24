package router

import(
	"github.com/gin-gonic/gin"
	"task-manager/controllers"
)

func InitializeRouter(){
	router := gin.Default()

	router.GET("/tasks",controllers.GetTasks)
	router.GET("/tasks/:id",controllers.GetSpecificTask)
	router.PUT("/tasks/:id",controllers.UpdateTask)
	router.POST("/tasks",controllers.CreateTask)
	router.DELETE("/tasks/:id",controllers.DeleteTask)

	router.Run("localhost:8080")
}