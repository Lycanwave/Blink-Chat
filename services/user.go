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

func CreateUser(user models.User) (*mongo.InsertOneResult, error) {

	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now
	inserted, err := db.GetCollection("blinkchat", "users").InsertOne(context.Background(), user)
	// defer inserted.Close()

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return inserted, nil
}

func GetUser(id string) (*models.User, error) {
	// Convert string ID to ObjectID
	_id, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid ObjectID: %v", err)
	}

	filter := bson.M{"_id": _id}

	var user models.User

	err = db.GetCollection("blinkchat", "users").FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		fmt.Println(err)
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("user not found")
		}

		return nil, fmt.Errorf("error finding user: %v", err)
	}

	return &user, nil
}

func GetUsers() ([]*models.User, error) {

	filter := bson.M{}

	curr, err := db.GetCollection("blinkchat", "users").Find(context.Background(), filter)

	if err != nil {
		return nil, err
	}

	var users []*models.User
	for curr.Next(context.Background()) {
		// var user models.User
		// var userBSON bson.M
		var user models.User
		err := curr.Decode(&user)

		if err != nil {
			return nil, err
		}

		users = append(users, &user)
	}

	// fmt.Println(cursor)
	return users, nil
}
