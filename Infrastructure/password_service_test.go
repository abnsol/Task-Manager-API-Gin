package infrastructure

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"golang.org/x/crypto/bcrypt"
)

// PasswordServiceSuite defines the test suite for PasswordService
type PasswordServiceSuite struct {
	suite.Suite
	passwordService PasswordService
}

// SetupTest runs before each test in the suite
func (s *PasswordServiceSuite) SetupTest() {
	s.passwordService = PasswordService{}
}

// TestPasswordServiceSuite runs the entire test suite
func TestPasswordServiceSuite(t *testing.T) {
	suite.Run(t, new(PasswordServiceSuite))
}

// TestHashPassword tests the HashPassword method
func (s *PasswordServiceSuite) TestHashPassword() {
	password := "my-secret-password"
	hashedPassword, err := s.passwordService.HashPassword(password)

	s.NoError(err)
	s.NotEmpty(hashedPassword)

	// Verify the hash corresponds to the password
	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	s.NoError(err)
}

// TestCheckPasswordCorrect tests checking a correct password
func (s *PasswordServiceSuite) TestCheckPasswordCorrect() {
	password := "my-secret-password"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	err := s.passwordService.CheckPassword(string(hashedPassword), password)
	s.NoError(err)
}

// TestCheckPasswordIncorrect tests checking an incorrect password
func (s *PasswordServiceSuite) TestCheckPasswordIncorrect() {
	password := "my-secret-password"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	err := s.passwordService.CheckPassword(string(hashedPassword), "wrong-password")
	s.Error(err)
	s.Equal(bcrypt.ErrMismatchedHashAndPassword, err)
}
