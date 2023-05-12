package services

import (
	entities "internal/entities"
	repository "internal/repository"
)

type UserService interface {
	CreateUser(user *entities.User) error
	UpdateUser(user *entities.User) error
	DeleteUser(id int) error
	GetUserByID(id int) (*entities.User, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (s *userService) CreateUser(user *entities.User) error {
	return s.userRepo.Create(user)
}

func (s *userService) UpdateUser(user *entities.User) error {
	return s.userRepo.Update(user)
}

func (s *userService) DeleteUser(id int) error {
	return s.userRepo.Delete(id)
}

func (s *userService) GetUserByID(id int) (*entities.User, error) {
	return s.userRepo.GetByID(id), nil
}
