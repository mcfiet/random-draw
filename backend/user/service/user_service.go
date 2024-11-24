package service

import (
	"errors"

	"github.com/google/uuid"
	"github.com/mcfiet/goDo/user/model"
	"github.com/mcfiet/goDo/user/repository"
	"github.com/mcfiet/goDo/utils"
)

type UserService struct {
	UserRepository repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo}
}

func (service *UserService) FindById(id uuid.UUID) (model.User, error) {
	return service.UserRepository.FindById(id)
}

func (service *UserService) FindByUsername(username string) (model.User, error) {
	return service.UserRepository.FindByUsername(username)
}

func (service *UserService) FindAll() ([]model.User, error) {
	return service.UserRepository.FindAll()
}

func (service *UserService) Save(user model.User) error {
	if user.Username == "" || user.Password == "" {
		return errors.New("Username, Email and Password are required")
	}
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashedPassword

	return service.UserRepository.Save(user)
}

func (service *UserService) Update(user model.User) error {
	return service.UserRepository.Update(user)
}

func (service *UserService) CheckIfUserExists(user model.User) error {
	_, err := service.UserRepository.FindByUsername(user.Username)
	if err == nil {
		return errors.New("User with username already exists")
	}

	return nil
}
