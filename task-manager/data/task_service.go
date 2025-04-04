package data

import (
	"context"
	"errors"
	"log"
	_"fmt"
	"task-manager/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	_"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// var tasks = make(map[int]models.Task)
// var nextID = 1



func GetAllTasks() []models.Task{
	var allTasks []models.Task

	cur,err := collection.Find(context.TODO(),bson.D{{"user_id",primitive.NilObjectID}})
	if err != nil{
		log.Fatal(err)
	}

	for cur.Next(context.TODO()){
		var task models.Task
		err := cur.Decode(&task)
		if err != nil{
			log.Fatal(err)
		}
		allTasks = append(allTasks,task)
	}
	cur.Close(context.TODO())
	
	return allTasks
}

func GetTaskByID(id primitive.ObjectID) (models.Task,error){
	var task models.Task
	filter := bson.D{primitive.E{Key:"_id",Value:id}}
	err := collection.FindOne(context.TODO(),filter).Decode(&task)
	if err != nil{
		if err == mongo.ErrNoDocuments{
			return models.Task{},errors.New("task not found")
		}
		return models.Task{},err
	}

	return task,nil
}

func UpdateTask(id primitive.ObjectID,newTask models.Task) (models.Task, error){
	filter := bson.D{primitive.E{Key:"_id",Value:id}}
	update := bson.D{primitive.E{Key:"$set",Value:newTask}}
	_,err := collection.UpdateOne(context.TODO(),filter,update)
	if err != nil{
		log.Fatal(err)
	}
	return newTask,nil
}

func CreateTask(newTask models.Task) models.Task{
	newTask.ID = primitive.NewObjectID()
	_,err := collection.InsertOne(context.TODO(),newTask)
	if err != nil{
		log.Fatal(err)
		return models.Task{}
	}
	return newTask
}

func DeleteTask(id primitive.ObjectID) error {

	filter := bson.M{"_id":id}
	result,err := collection.DeleteOne(context.TODO(),filter)
	if err != nil{
		return err
	}
	if result.DeletedCount == 0{
		return errors.New("task not found")
	}
	return nil 
}

