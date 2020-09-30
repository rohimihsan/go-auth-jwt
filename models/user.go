package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Firstname  string             `json:"firstname" validate:"required"`
	Lastname   string             `json:"lastname" validate:"required"`
	Username   string             `json:"username"`
	Email      string             `json:"email" validate:"required,email"`
	Password   string             `validate:"required"`
	Created_at time.Time          `json:"created_at"`
	Updated_at time.Time          `json:"updated_at"`
}
