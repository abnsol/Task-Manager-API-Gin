package data

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Db *mongo.Database
var UserCollection *mongo.Collection
var TaskCollection *mongo.Collection

func InitDB() {
	mongoURI := os.Getenv("mongoURI")
	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	Db = client.Database("Tasks-DataBase")
	UserCollection = Db.Collection("Users")
	TaskCollection = Db.Collection("Tasks")
}
