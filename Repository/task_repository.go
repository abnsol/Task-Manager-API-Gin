package repository

import (
	"context"
	"errors"
	"log"
	domain "task_management/Domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TaskRepository struct {
	database   mongo.Database
	collection string
}

var errTask = errors.New("task not found")

func NewTaskRepository(db mongo.Database, collection string) domain.TaskRepository {
	return &TaskRepository{
		database:   db,
		collection: collection,
	}
}

func (tr *TaskRepository) GetTasks() []domain.Task {
	taskCollection := tr.database.Collection(tr.collection)

	// cursor
	findOptions := options.Find()

	// Here's an array in which you can store the decoded documents
	var tasks []domain.Task

	// Passing bson.D{} as the filter matches all documents in the data.taskCollectioncollection
	cur, err := taskCollection.Find(context.TODO(), bson.D{}, findOptions)
	if err != nil {
		panic(err)
	}

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem domain.Task
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

func (tr *TaskRepository) CreateTask(task domain.Task) (string, error) {
	taskCollection := tr.database.Collection(tr.collection)
	_, err := taskCollection.InsertOne(context.TODO(), task)

	if err != nil {
		return "", err
	}
	return "Task Created Successfully", nil

}

func (tr *TaskRepository) GetTaskById(id string) (domain.Task, error) {
	taskCollection := tr.database.Collection(tr.collection)
	var task domain.Task

	err := taskCollection.FindOne(context.TODO(), bson.M{"id": id}).Decode(&task)
	if err != nil {
		return domain.Task{}, errTask
	}

	return task, nil

}

func (tr *TaskRepository) ReplaceTask(id string, newTask domain.Task) (domain.Task, error) {
	taskCollection := tr.database.Collection(tr.collection)
	filter := bson.D{{Key: "id", Value: id}}
	res, _ := taskCollection.ReplaceOne(context.TODO(), filter, newTask)

	if res.MatchedCount == 0 {
		return domain.Task{}, errors.New("id not found")
	}

	return newTask, nil

}

func (tr *TaskRepository) DeleteTask(id string) (string, error) {
	taskCollection := tr.database.Collection(tr.collection)

	result, err := taskCollection.DeleteOne(context.TODO(), bson.M{"id": id})
	if err != nil {
		return "", err
	}
	if result.DeletedCount == 0 {
		return "", errTask
	}
	return "Task Deleted Successfully", nil

}
