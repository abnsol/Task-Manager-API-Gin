package data

// Import the bcrypt package
import (
	"context"
	"errors"
	"log"
	"os"
	"task_management/models"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func Register(user models.User) (string, error) {
	// Check if user already exists
	var result models.User
	err := UserCollection.FindOne(context.TODO(), bson.M{"email": user.Email}).Decode(&result)

	if err == nil {
		return "", errors.New("user with this email already exists")
	}

	// User registration logic
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New("internal server error")
	}

	user.Password = string(hashedPassword)

	// Generate a new ID for the user
	user.ID = primitive.NewObjectID()

	_, err = UserCollection.InsertOne(context.TODO(), user)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	return "user successfully registered", nil
}

func Login(user models.User) (string, string, error) {
	var jwtSecret = []byte(os.Getenv("jwt_secret"))
	// Check if user exists
	var existingUser models.User
	err := UserCollection.FindOne(context.TODO(), bson.M{"email": user.Email}).Decode(&existingUser)
	if err != nil || bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password)) != nil {
		return "", "", errors.New("invalid email or password")
	}

	// Generate JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": existingUser.ID,
		"email":   existingUser.Email,
		"role":    existingUser.Role,
	})

	jwtToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", "", errors.New("internal server error")
	}

	return "User logged in successfully", jwtToken, nil
}

func UpdateUserRole(email string, newRole string) error {
	filter := bson.M{"email": email}
	update := bson.M{"$set": bson.M{"role": newRole}}

	result, err := UserCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return errors.New("user not found")
	}
	return nil
}
