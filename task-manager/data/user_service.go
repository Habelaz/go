package data

import (
	"context"
	"task-manager/models"
	"errors"
	"log"
	"time"
	"os"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)
func init(){
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

var JwtSecret = []byte(os.Getenv("JWT_SECRET"))

func UserRegister(newUser models.User) (models.User,error){
	// check if user already exists
	filter := bson.D{primitive.E{Key:"username",Value:newUser.Username}}
	var existingUser models.User
	err := userCollection.FindOne(context.TODO(),filter).Decode(&existingUser)
	if err == nil{
		return models.User{},errors.New("user already exists")
	}
	
	// creating new user if user does not exist
	newUser.ID = primitive.NewObjectID()
	hashedPassword,err := bcrypt.GenerateFromPassword([]byte(newUser.Password),bcrypt.DefaultCost)
	if err != nil{
		log.Fatal(err)
		return models.User{},errors.New("error creating user")
	}
	newUser.Password = string(hashedPassword)
	
	
	_, err = userCollection.InsertOne(context.TODO(), newUser)
	
	if err != nil{
		log.Fatal(err)
		return models.User{},errors.New("error creating user")
	}

	return newUser,nil

}

func UserLogin(user models.User) (models.User, error, string) {
	filter := bson.D{primitive.E{Key: "username", Value: user.Username}}
	var existingUser models.User

	err := userCollection.FindOne(context.TODO(), filter).Decode(&existingUser)
	if err != nil {
		fmt.Println("FindOne Error:", err) 
		return models.User{}, errors.New("user not found"), ""
	}

	// Compare hashed passwords
	err = bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password))
	if err != nil {
		fmt.Println("Password Error:", err) 
		return models.User{}, errors.New("invalid username or password"), ""
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": existingUser.Username,
		"role":     existingUser.Role,
		"id":       existingUser.ID.Hex(),
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	})

	jwtToken, err := token.SignedString([]byte(JwtSecret))
	if err != nil {
		fmt.Println("JWT Error:", err) // Debugging line
		return models.User{}, errors.New("error generating token"), ""
	}

	return existingUser, nil, jwtToken
}
