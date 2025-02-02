package services

import (
	"context"
	"fmt"
	"sync"

	"go-lang/blinkchat/models"
	db "go-lang/blinkchat/utils"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var wg sync.WaitGroup

func CreateMessage(message *models.Message) (*mongo.InsertOneResult, error) {

	now := time.Now()
	message.CreatedAt = now
	message.UpdatedAt = now
	inserted, err := db.GetCollection("blinkchat", "message").InsertOne(context.Background(), message)
	// defer inserted.Close()

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return inserted, nil
}

func GetMessage(id string) (*models.Message, error) {

	conversation_id, _ := bson.ObjectIDFromHex(id)
	filter := bson.M{"chat_id": conversation_id}
	var message *models.Message

	err := db.GetCollection("blinkchat", "message").FindOne(context.Background(), filter).Decode(&message)

	if err != nil {
		return nil, err
	}

	return message, nil
}

func dbCall(filter bson.M) (*mongo.Cursor, error) {

	data, err := db.GetCollection("blinkchat", "message").Find(context.Background(), filter)
	return data, err
}

func GetMessages(id string) ([]*models.Message, error) {
	fmt.Println("entered")
	_id, _ := bson.ObjectIDFromHex(id)
	filter := bson.M{"chat_id": _id}

	var curr *mongo.Cursor
	var err error = nil

	wg.Add(1)
	go func() {
		defer wg.Done()
		curr, err = dbCall(filter) // Execute the database call
	}()

	wg.Wait()

	if err != nil {
		return nil, err
	}

	var messages []*models.Message
	for curr.Next(context.Background()) {
		// var user models.User
		// var userBSON bson.M
		var message models.Message
		err := curr.Decode(&message)

		if err != nil {
			return nil, err
		}

		messages = append(messages, &message)
	}

	// fmt.Println(cursor)
	return messages, nil
}
