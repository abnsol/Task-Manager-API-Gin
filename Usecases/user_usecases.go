package usecases

import domain "task_management/Domain"

type UserUseCase struct {
	UserRepository domain.UserRepository
}

func NewUserUseCase(userRepository domain.UserRepository) domain.UserUseCase {
	return &UserUseCase{
		UserRepository: userRepository,
	}
}

func (uc *UserUseCase) Register(user domain.User) (string, error) {
	return uc.UserRepository.Register(user)
}

func (uc *UserUseCase) Login(user domain.User) (string, string, error) {
	return uc.UserRepository.Login(user)
}

func (uc *UserUseCase) UpdateUserRole(email string, newRole string) error {
	return uc.UserRepository.UpdateUserRole(email, newRole)
}
