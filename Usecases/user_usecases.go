package usecases

import (
	"errors"
	domain "task_management/Domain"
)

type UserUseCase struct {
	UserRepository  domain.IUserRepository
	JwtService      domain.IJwtService
	PasswordService domain.IPasswordService
}

func NewUserUseCase(userRepository domain.IUserRepository, passwordService domain.IPasswordService, jwtService domain.IJwtService) domain.IUserUseCase {
	return &UserUseCase{
		UserRepository:  userRepository,
		JwtService:      jwtService,
		PasswordService: passwordService,
	}
}

func (uc *UserUseCase) Register(user domain.User) (string, error) {
	// Check if user already exists
	checkUser, _ := uc.UserRepository.CheckUserExists(user)

	if checkUser {
		return "", errors.New("user with this email already exists")
	}

	// User registration logic
	hashedPassword, err := uc.PasswordService.HashPassword(user.Password)
	if err != nil {
		return "", errors.New("error hashing password")
	}

	user.Password = string(hashedPassword)
	return uc.UserRepository.Register(user)
}

func (uc *UserUseCase) Login(user domain.User) (string, string, error) {
	checkUser, existingUser := uc.UserRepository.CheckUserExists(user)

	if !checkUser {
		return "", "", errors.New("email address doesn't exist")
	}

	if uc.PasswordService.CheckPassword(existingUser.Password, user.Password) != nil {
		return "", "", errors.New("password incorrect")
	}

	jwtToken, err := uc.JwtService.GenerateToken(existingUser.ID, existingUser.Email, existingUser.Role)
	if err != nil {
		return "", "", errors.New("error generating token")
	}

	return "User logged in successfully", jwtToken, nil
}

func (uc *UserUseCase) UpdateUserRole(email string, newRole string) error {
	return uc.UserRepository.UpdateUserRole(email, newRole)
}
