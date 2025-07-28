package repository

import (
	"context"
	"errors"
	"log"
	domain "task_management/Domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	database   mongo.Database
	collection string
}

func NewUserRepository(db mongo.Database, collection string) domain.IUserRepository {
	return &UserRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *UserRepository) Register(user domain.User) (string, error) {
	userCollection := ur.database.Collection(ur.collection)

	// Generate a new ID for the user
	user.ID = primitive.NewObjectID()
	_, err := userCollection.InsertOne(context.TODO(), user)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	return "user successfully registered", nil
}

func (ur *UserRepository) UpdateUserRole(email string, newRole string) error {
	userCollection := ur.database.Collection(ur.collection)
	filter := bson.M{"email": email}
	update := bson.M{"$set": bson.M{"role": newRole}}

	result, err := userCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return errors.New("user not found")
	}
	return nil
}

func (ur *UserRepository) CheckUserExists(user domain.User) (bool, domain.User) {
	var result domain.User
	userCollection := ur.database.Collection(ur.collection)

	err := userCollection.FindOne(context.TODO(), bson.M{"email": user.Email}).Decode(&result)
	return err == nil, result
}
