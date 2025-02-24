package data

import (
	"errors"
	"task-manager/models"
)

var tasks = make(map[int]models.Task)
var nextID = 1

func GetAllTasks() []models.Task{
	var allTasks []models.Task
	for _,task := range tasks {
		allTasks = append(allTasks,task)
	}
	return allTasks
}

func GetTaskByID(id int) (models.Task,error){
	task,exist := tasks[id]
	if !exist{
		return models.Task{} , errors.New("task no found")
	}
	return task,nil
}

func UpdateTask(id int,newTask models.Task) (models.Task, error){
	_,exist := tasks[id]
	if !exist{
		return models.Task{},errors.New("task not found")
	}
	newTask.ID = id
	tasks[id] = newTask
	return newTask,nil
}

func CreateTask(newTask models.Task) models.Task{
	newTask.ID = nextID
	tasks[nextID] = newTask
	nextID += 1
	return newTask
}

func DeleteTask(id int) error {
	_,exist := tasks[id]
	if !exist{
		return errors.New("task not found")
	}
	delete(tasks,id)
	return nil
}

