package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Member struct {
	UserId bson.ObjectID `json:"user_id" bson:"user_id"`
	Role   string        `json:"role" bson:"role"`
}

type Chat struct {
	Id        bson.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Type      string        `json:"type,omitempty" bson:"type,omitempty"`
	Name      string        `json:"name,omitempty" bson:"name,omitempty"`
	Members   []*Member     `json:"members,omitempty" bson:"members,omitempty"`
	CreatedBy bson.ObjectID `json:"created_by,omitempty" bson:"created_by,omitempty"`
	CreatedAt time.Time     `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time     `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
