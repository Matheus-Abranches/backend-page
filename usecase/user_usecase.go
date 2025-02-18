package usecase

import (
	"my-project/model"
	"my-project/repository"
)

type UserUseCase struct {
	repository repository.UserRepository
}

func NewUserUseCase(repo repository.UserRepository) UserUseCase {
	return UserUseCase{
		repository: repo,
	}
}

func (uc *UserUseCase) GetUsers() ([]model.User, error) {
	return uc.repository.GetUsers()
}

func (uc *UserUseCase) CreateUser(user model.User) error {
	return uc.repository.CreateUser(user)
}

func (uc *UserUseCase) UpdateUser(id uint, user model.User) error {
	return uc.repository.UpdateUser(id, user)
}

func (uc *UserUseCase) DeleteUser(id string) error {
	return uc.repository.DeleteUser(id)
}
