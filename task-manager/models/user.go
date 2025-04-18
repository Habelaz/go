package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct{
	ID primitive.ObjectID `bson:"_id,omitempty"`
	Username string `json:"username",bson:"username omitempty"`
	Password string `json:"password",bson:"password omitempty"`
	Role string `json:"role",bson:"role omitempty"`
}