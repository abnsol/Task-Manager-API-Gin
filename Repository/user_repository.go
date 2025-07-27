package repository

import (
	"context"
	"errors"
	"log"
	domain "task_management/Domain"
	infrastructure "task_management/Infrastructure"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	database   mongo.Database
	collection string
}

func NewUserRepository(db mongo.Database, collection string) domain.UserRepository {
	return &UserRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *UserRepository) Register(user domain.User) (string, error) {
	userCollection := ur.database.Collection(ur.collection)

	// Check if user already exists
	var result domain.User
	err := userCollection.FindOne(context.TODO(), bson.M{"email": user.Email}).Decode(&result)

	if err == nil {
		return "", errors.New("user with this email already exists")
	}

	// User registration logic
	hashedPassword, err := infrastructure.HashPassword(user.Password)
	if err != nil {
		return "", errors.New("internal server error")
	}

	user.Password = string(hashedPassword)

	// Generate a new ID for the user
	user.ID = primitive.NewObjectID()

	_, err = userCollection.InsertOne(context.TODO(), user)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	return "user successfully registered", nil

}
func (ur *UserRepository) Login(user domain.User) (string, string, error) {
	userCollection := ur.database.Collection(ur.collection)
	// var jwtSecret = []byte(os.Getenv("jwt_secret"))
	// Check if user exists
	var existingUser domain.User
	err := userCollection.FindOne(context.TODO(), bson.M{"email": user.Email}).Decode(&existingUser)
	if err != nil || infrastructure.CheckPassword(existingUser.Password, user.Password) != nil {
		return "", "", errors.New("invalid email or password")
	}

	// // Generate JWT
	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	// 	"user_id": existingUser.ID,
	// 	"email":   existingUser.Email,
	// 	"role":    existingUser.Role,
	// })

	jwtToken, err := infrastructure.GenerateToken(existingUser.ID, existingUser.Email, existingUser.Role)
	if err != nil {
		return "", "", errors.New("internal server error")
	}

	return "User logged in successfully", jwtToken, nil

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
