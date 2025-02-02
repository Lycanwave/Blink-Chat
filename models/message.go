package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Attachment struct {
	Type string `json:"type" bson:"type"`
	URL  string `json:"url" bson:"url"`
}

type Message struct {
	Id          bson.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	ChatId      bson.ObjectID `json:"chat_id,omitempty" bson:"chat_id,omitempty"`
	SenderId    bson.ObjectID `json:"sender_id,omitempty" bson:"sender_id,omitempty"`
	Content     *string       `json:"content,omitempty" bson:"content,omitempty"`
	Attachments []*Attachment `json:"attachments,omitempty" bson:"attachments,omitempty"`
	IsRead      bool          `json:"is_read,omitempty" bson:"is_read,omitempty"`
	CreatedAt   time.Time     `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt   time.Time     `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
