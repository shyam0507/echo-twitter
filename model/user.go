package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Email     string             `json:"email" bson:"email"`
	Password  string             `json:"password,omitempty" bson:"password"`
	Token     string             `json:"token,omitempty" bson:"-"`
	Followers []string           `json:"followers,omitempty" bson:"followers,omitempty"`
}
