package infrastructure

import (
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JwtService struct{}

func (JwtService) GenerateToken(user_id primitive.ObjectID, userEmail string, userRole string) (string, error) {
	jwtSecret := []byte(os.Getenv("jwt_secret"))

	// Generate JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user_id,
		"email":   userEmail,
		"role":    userRole,
	})

	jwtToken, err := token.SignedString(jwtSecret)
	return jwtToken, err
}

func (JwtService) ValidateToken(tokenString string) (*jwt.Token, error) {
	jwtSecret := []byte(os.Getenv("jwt_secret"))
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})
}
