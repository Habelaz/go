package controllers

import (
	"net/http"
	_"strconv"
	"task-manager/data"
	"task-manager/models"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gin-gonic/gin"
)

func GetTasks(c *gin.Context) {
	tasks := data.GetAllTasks()
	c.IndentedJSON(http.StatusOK, tasks)

}

func GetSpecificTask(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return
	}
	task, err := data.GetTaskByID(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, task)

}

func UpdateTask(c *gin.Context) {
	var updatedTask models.Task
	if err := c.BindJSON(&updatedTask); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON"})
		return
	}
	
	objectID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return
	}
	if _, err := data.UpdateTask(objectID, updatedTask); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Task updated successfully"})
}

func CreateTask(c *gin.Context) {
	var newTask models.Task
	if err := c.BindJSON(&newTask); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON"})
		return
	} 
	
	data.CreateTask(newTask)

	c.IndentedJSON(http.StatusCreated,newTask)
}

func DeleteTask(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
	}
	data.DeleteTask(id)
	c.IndentedJSON(http.StatusOK,gin.H{"message":"Deleted Succesfully"})
}

func Register(c *gin.Context){
	var user models.User
	if err := c.ShouldBindBodyWithJSON(&user); err != nil{
		c.JSON(http.StatusBadRequest,err)
	}
	if user.Role == ""{
		user.Role = "user"
	}
	newUser,err := data.UserRegister(user)
	if err != nil{
		c.JSON(http.StatusInternalServerError,err)
		return
	}
	c.JSON(http.StatusCreated,newUser)

}

func Login(c *gin.Context){
	var user models.User
	if err := c.ShouldBindBodyWithJSON(&user); err != nil{
		c.JSON(http.StatusBadRequest,err)
	}
	existingUser,err,token := data.UserLogin(user)
	if err !=nil {
		c.JSON(http.StatusInternalServerError,err)
		return
	}
	c.JSON(http.StatusOK,gin.H{"token":token,"user":existingUser})
}