package models

//This file contains details or entries that will be created in db

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Entry struct {
	ID          primitive.ObjectID `bson:"id"`
	Dish        string             `json:"dish"`
	Fat         float64            `json:"fat"`
	Ingredients string             `json:"ingredient"`
	Calories    string             `json:"calories"`
}
