package main

import (
	"task-manager/data"
	"task-manager/router"
	"os"
	"log"
)

func main(){

	mongodb_uri := os.Getenv("MONGODB_URI")
	if mongodb_uri == ""{
		log.Fatal("Set your 'MONGODB_URI' environment variable.")
	}
	data.InitMongodb(mongodb_uri)
	router.InitializeRouter()
	defer data.DisconnectMongo()

}