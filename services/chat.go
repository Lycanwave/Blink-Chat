package services

import (
	"context"
	"fmt"

	"go-lang/blinkchat/models"
	db "go-lang/blinkchat/utils"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func CreateChat(chat models.Chat) (*mongo.InsertOneResult, error) {

	now := time.Now()
	chat.CreatedAt = now
	chat.UpdatedAt = now
	inserted, err := db.GetCollection("blinkchat", "chats").InsertOne(context.Background(), chat)
	// defer inserted.Close()

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return inserted, nil
}

func GetChat(id string) (*models.Chat, error) {

	_id, _ := bson.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id}
	var chat *models.Chat

	err := db.GetCollection("blinkchat", "chats").FindOne(context.Background(), filter).Decode(&chat)
	fmt.Println(chat)
	if err != nil {
		return nil, err
	}

	return chat, nil
}
