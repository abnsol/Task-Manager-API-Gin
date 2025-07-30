package usecases

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"

	domain "task_management/Domain"
	mock_domain "task_management/Mocks/domain"
)

// UserUseCaseSuite defines the test suite
type UserUseCaseSuite struct {
	suite.Suite
	mockUserRepo *mock_domain.IUserRepository
	mockPassSvc  *mock_domain.IPasswordService
	mockJwtSvc   *mock_domain.IJwtService
	userUseCase  domain.IUserUseCase
	user         domain.User
	existingUser domain.User
}

// SetupTest is run before each test in the suite
func (s *UserUseCaseSuite) SetupTest() {
	// Initialize mocks
	s.mockUserRepo = new(mock_domain.IUserRepository)
	s.mockPassSvc = new(mock_domain.IPasswordService)
	s.mockJwtSvc = new(mock_domain.IJwtService)

	// Initialize the use case with the mocks
	s.userUseCase = NewUserUseCase(s.mockUserRepo, s.mockPassSvc, s.mockJwtSvc)

	// Define common test data
	s.user = domain.User{
		Email:    "test@example.com",
		Password: "password123",
	}
	s.existingUser = domain.User{
		ID:       primitive.NewObjectID(),
		Email:    "test@example.com",
		Password: "hashedpassword",
		Role:     "user",
	}
}

// TestUserUseCaseSuite runs the entire test suite
func TestUserUseCaseSuite(t *testing.T) {
	suite.Run(t, new(UserUseCaseSuite))
}

// TestRegisterSuccess tests the successful registration scenario
func (s *UserUseCaseSuite) TestRegisterSuccess() {
	// Setup expectations
	s.mockUserRepo.On("CheckUserExists", s.user).Return(false, domain.User{}).Once()
	s.mockPassSvc.On("HashPassword", s.user.Password).Return([]byte("hashedpassword"), nil).Once()
	s.mockUserRepo.On("Register", mock.AnythingOfType("domain.User")).Return("some-user-id", nil).Once()

	// Execute the method
	id, err := s.userUseCase.Register(s.user)

	// Assertions
	s.NoError(err)
	s.Equal("some-user-id", id)
	s.mockUserRepo.AssertExpectations(s.T())
	s.mockPassSvc.AssertExpectations(s.T())
}

// TestRegisterUserAlreadyExists tests registration when a user already exists
func (s *UserUseCaseSuite) TestRegisterUserAlreadyExists() {
	s.mockUserRepo.On("CheckUserExists", s.user).Return(true, domain.User{}).Once()

	_, err := s.userUseCase.Register(s.user)

	s.Error(err)
	s.Equal("user with this email already exists", err.Error())
	s.mockUserRepo.AssertExpectations(s.T())
}

// TestRegisterErrorHashingPassword tests registration with a password hashing error
func (s *UserUseCaseSuite) TestRegisterErrorHashingPassword() {
	s.mockUserRepo.On("CheckUserExists", s.user).Return(false, domain.User{}).Once()
	s.mockPassSvc.On("HashPassword", s.user.Password).Return(nil, errors.New("hash error")).Once()

	_, err := s.userUseCase.Register(s.user)

	s.Error(err)
	s.Equal("error hashing password", err.Error())
	s.mockUserRepo.AssertExpectations(s.T())
	s.mockPassSvc.AssertExpectations(s.T())
}

// TestLoginSuccess tests the successful login scenario
func (s *UserUseCaseSuite) TestLoginSuccess() {
	s.mockUserRepo.On("CheckUserExists", s.user).Return(true, s.existingUser).Once()
	s.mockPassSvc.On("CheckPassword", s.existingUser.Password, s.user.Password).Return(nil).Once()
	s.mockJwtSvc.On("GenerateToken", s.existingUser.ID, s.existingUser.Email, s.existingUser.Role).Return("test-token", nil).Once()

	msg, token, err := s.userUseCase.Login(s.user)

	s.NoError(err)
	s.Equal("User logged in successfully", msg)
	s.Equal("test-token", token)
	s.mockUserRepo.AssertExpectations(s.T())
	s.mockPassSvc.AssertExpectations(s.T())
	s.mockJwtSvc.AssertExpectations(s.T())
}

// TestLoginEmailDoesNotExist tests login when an email does not exist
func (s *UserUseCaseSuite) TestLoginEmailDoesNotExist() {
	s.mockUserRepo.On("CheckUserExists", s.user).Return(false, domain.User{}).Once()

	_, _, err := s.userUseCase.Login(s.user)

	s.Error(err)
	s.Equal("email address doesn't exist", err.Error())
	s.mockUserRepo.AssertExpectations(s.T())
}

// TestLoginPasswordIncorrect tests login with an incorrect password
func (s *UserUseCaseSuite) TestLoginPasswordIncorrect() {
	s.mockUserRepo.On("CheckUserExists", s.user).Return(true, s.existingUser).Once()
	s.mockPassSvc.On("CheckPassword", s.existingUser.Password, s.user.Password).Return(errors.New("password mismatch")).Once()

	_, _, err := s.userUseCase.Login(s.user)

	s.Error(err)
	s.Equal("password incorrect", err.Error())
	s.mockUserRepo.AssertExpectations(s.T())
	s.mockPassSvc.AssertExpectations(s.T())
}

// TestLoginErrorGeneratingToken tests login with a token generation error
func (s *UserUseCaseSuite) TestLoginErrorGeneratingToken() {
	s.mockUserRepo.On("CheckUserExists", s.user).Return(true, s.existingUser).Once()
	s.mockPassSvc.On("CheckPassword", s.existingUser.Password, s.user.Password).Return(nil).Once()
	s.mockJwtSvc.On("GenerateToken", s.existingUser.ID, s.existingUser.Email, s.existingUser.Role).Return("", errors.New("jwt error")).Once()

	_, _, err := s.userUseCase.Login(s.user)

	s.Error(err)
	s.Equal("error generating token", err.Error())
	s.mockUserRepo.AssertExpectations(s.T())
	s.mockPassSvc.AssertExpectations(s.T())
	s.mockJwtSvc.AssertExpectations(s.T())
}

// TestUpdateUserRoleSuccess tests successful user role update
func (s *UserUseCaseSuite) TestUpdateUserRoleSuccess() {
	email := "test@example.com"
	newRole := "admin"
	s.mockUserRepo.On("UpdateUserRole", email, newRole).Return(nil).Once()

	err := s.userUseCase.UpdateUserRole(email, newRole)

	s.NoError(err)
	s.mockUserRepo.AssertExpectations(s.T())
}

// TestUpdateUserRoleFailure tests user role update failure
func (s *UserUseCaseSuite) TestUpdateUserRoleFailure() {
	email := "test@example.com"
	newRole := "admin"
	s.mockUserRepo.On("UpdateUserRole", email, newRole).Return(errors.New("db error")).Once()

	err := s.userUseCase.UpdateUserRole(email, newRole)

	s.Error(err)
	s.mockUserRepo.AssertExpectations(s.T())
}
