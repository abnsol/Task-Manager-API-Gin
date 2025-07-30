package infrastructure

import (
	"os"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JwtServiceSuite struct {
	suite.Suite
	jwtService     JwtService
	originalSecret string
}

func (s *JwtServiceSuite) SetupTest() {
	s.jwtService = JwtService{}
	s.originalSecret = os.Getenv("jwt_secret")
	os.Setenv("jwt_secret", "a-very-secure-and-not-at-all-hardcoded-secret")
}

// TearDownTest runs after each test in the suite.
// It restores the original environment variable.
func (s *JwtServiceSuite) TearDownTest() {
	os.Setenv("jwt_secret", s.originalSecret)
}

// TestJwtServiceSuite runs the entire test suite
func TestJwtServiceSuite(t *testing.T) {
	suite.Run(t, new(JwtServiceSuite))
}

// TestGenerateToken tests the GenerateToken method
func (s *JwtServiceSuite) TestGenerateToken() {
	userID := primitive.NewObjectID()
	email := "test@example.com"
	role := "user"

	tokenString, err := s.jwtService.GenerateToken(userID, email, role)

	s.NoError(err)
	s.NotEmpty(tokenString)

	// Parse the token to verify its claims
	token, err := s.jwtService.ValidateToken(tokenString)

	s.NoError(err)
	s.True(token.Valid)

	claims, ok := token.Claims.(jwt.MapClaims)
	s.True(ok)
	s.Equal(userID.Hex(), claims["user_id"])
	s.Equal(email, claims["email"])
	s.Equal(role, claims["role"])
}

// TestValidateToken tests the ValidateToken method
func (s *JwtServiceSuite) TestValidateToken() {
	userID := primitive.NewObjectID()
	secret := []byte(os.Getenv("jwt_secret"))

	// Create various tokens for testing
	validToken, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 1).Unix(),
	}).SignedString(secret)

	expiredToken, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(-time.Hour * 1).Unix(),
	}).SignedString(secret)

	invalidSignatureToken, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
	}).SignedString([]byte("wrong-secret"))

	s.Run("valid token", func() {
		token, err := s.jwtService.ValidateToken(validToken)
		s.NoError(err)
		s.True(token.Valid)
	})

	s.Run("expired token", func() {
		_, err := s.jwtService.ValidateToken(expiredToken)
		s.Error(err)
	})

	s.Run("invalid signature", func() {
		_, err := s.jwtService.ValidateToken(invalidSignatureToken)
		s.Error(err)
	})
}
