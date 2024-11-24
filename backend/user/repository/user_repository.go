package repository

import (
	"github.com/google/uuid"
	"github.com/mcfiet/goDo/user/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindById(id uuid.UUID) (model.User, error)
	FindByUsername(username string) (model.User, error)
	FindByEmail(email string) (model.User, error)
	FindByUsernameOrEmail(username string, email string) (model.User, error)
	FindAll() ([]model.User, error)
	Save(user model.User) error
	Update(user model.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (repo *userRepository) FindById(id uuid.UUID) (model.User, error) {
	var user model.User
	result := repo.db.First(&user, id)
	return user, result.Error
}

func (repo *userRepository) FindByUsername(username string) (model.User, error) {
	var user model.User
	result := repo.db.Where("username = ?", username).First(&user)
	return user, result.Error
}

func (repo *userRepository) FindByEmail(email string) (model.User, error) {
	var user model.User
	result := repo.db.Where("email = ?", email).First(&user)

	return user, result.Error
}

func (repo *userRepository) FindByUsernameOrEmail(username string, email string) (model.User, error) {
	var user model.User
	result := repo.db.Where("username = ? OR email = ?", username, email).First(&user)

	return user, result.Error
}

func (repo *userRepository) FindAll() ([]model.User, error) {
	var users []model.User
	result := repo.db.Find(&users)
	return users, result.Error
}

func (repo *userRepository) Save(user model.User) error {
	result := repo.db.Create(&user)
	return result.Error
}

func (repo *userRepository) Update(user model.User) error {
	result := repo.db.Model(&user).Updates(user)
	return result.Error
}
