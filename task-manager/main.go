package main

import (
	
	"log"
	"os"
	"task-manager/data"
	"task-manager/router"
	"github.com/joho/godotenv"
)


func main(){
	err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    mongodb_uri := os.Getenv("MONGODB_URI")

	if mongodb_uri == ""{
		log.Fatal("Set your 'MONGODB_URI' environment variable.")
	}
	data.InitMongodb(mongodb_uri)
	router.InitializeRouter()
	defer data.DisconnectMongo()

}