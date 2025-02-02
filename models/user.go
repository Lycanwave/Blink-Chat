package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type User struct {
	Id         bson.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserName   string        `json:"user_name" bson:"user_name"`
	PhoneNo    string        `json:"phone_no" bson:"phone_no"`
	Email      string        `json:"email" bson:"email"`
	Password   string        `json:"password" bson:"password"`
	Status     string        `json:"status" bson:"status"`
	ProfileUrl string        `json:"profile_url,omitempty" bson:"profile_url,omitempty"`
	LastSeen   time.Time     `json:"last_seen,omitempty" bson:"last_seen,omitempty"`
	CreatedAt  time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt  time.Time     `json:"updated_at" bson:"updated_at"`
}
