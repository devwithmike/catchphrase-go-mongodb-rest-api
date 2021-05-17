package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Catchphrase struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	MovieName    string             `json:"movieName,omitempty" bson:"movieName,omitempty"`
	Catchphrase  string             `json:"catchphrase,omitempty" bson:"catchphrase,omitempty"`
	MovieContext string             `json:"movieContext,omitempty" bson:"movieContext,omitempty"`
}
