package data

import (
	"context"
	"errors"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"task_management/models"
)

// connect to db
var db = connect_db()
var collection = db.Database("Tasks-DataBase").Collection("Tasks")

var errTask = errors.New("task not found")

// get all tasks
func GetTasks() []models.Task {
	// cursor
	findOptions := options.Find()

	// Here's an array in which you can store the decoded documents
	var tasks []models.Task

	// Passing bson.D{} as the filter matches all documents in the collection
	cur, err := collection.Find(context.TODO(), bson.D{}, findOptions)
	if err != nil {
		panic(err)
	}

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem models.Task
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		tasks = append(tasks, elem)
	}

	if err := cur.Err(); err != nil {
		panic(err)
	}

	// Close the cursor once finished
	cur.Close(context.TODO())
	return tasks
}

// returns error if task not found
func GetTaskById(id string) (models.Task, error) {
	var task models.Task

	err := collection.FindOne(context.TODO(), bson.M{"id": id}).Decode(&task)
	if err != nil {
		return models.Task{}, errTask
	}

	return task, nil
}

func ReplaceTask(id string, newTask models.Task) (models.Task, error) {
	filter := bson.D{{Key: "id", Value: id}}
	res, _ := collection.ReplaceOne(context.TODO(), filter, newTask)

	if res.MatchedCount == 0 {
		return models.Task{}, errors.New("id not found")
	}

	return newTask, nil
}

func DeleteTask(id string) (string, error) {
	result, err := collection.DeleteOne(context.TODO(), bson.M{"id": id})
	if err != nil {
		return "", err
	}
	if result.DeletedCount == 0 {
		return "", errTask
	}
	return "Task Deleted Successfully", nil
}

// always creates Task so no need for error
func CreateTask(task models.Task) (string, error) {
	_, err := collection.InsertOne(context.TODO(), task)

	if err != nil {
		return "", err
	}
	return "Task Created Successfully", nil
}

func connect_db() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb+srv://<username>:<password>@cluster0.mongodb.net/?retryWrites=true&w=majority")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		panic(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to MongoDB!")
	return client
}
