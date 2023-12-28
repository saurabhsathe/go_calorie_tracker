package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Entry struct {
	Id          primitive.ObjectID `bson:"id"`
	Dish        string             `json:"dish"`
	Ingredients string             `json:"ingredients"`
	Fat         float64            `json:"fat"`
	Calories    float64            `json:"calories"`
}
