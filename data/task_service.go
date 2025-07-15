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

// In-memory storage for tasks
// var tasks = []models.Task{
// 	{ID: "1", Title: "Go and Gin", Description: "Finish RestAPI basics with go and gin", Time: "12:30", Status: false},
// 	{ID: "2", Title: "PostMan", Description: "Start Project and Write Docs with post man", Time: "3:30", Status: false},
// 	{ID: "3", Title: "Mongo DB", Description: "Add persistency to the project", Time: "6:30", Status: false},
// }

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

// get task Index
// func getIndex(newTask models.Task) (int, error) {
// 	for idx, task := range tasks {
// 		if newTask.ID == task.ID {
// 			return idx, nil
// 		}
// 	}

// 	return -1, errors.New("index not found")
// }

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
	clientOptions := options.Client().ApplyURI("mongodb+srv://abensol:h92JAgiSLsdLezck@cluster0.i8xepli.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0")

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
