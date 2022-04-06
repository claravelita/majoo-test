package postgres

import (
	"errors"
	"github.com/claravelita/majoo-test/domain"
	"github.com/claravelita/majoo-test/repository"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserReposiory {
	return userRepository{
		db,
	}
}

func (r userRepository) FetchUserByUsername(username string) (*domain.User, error) {
	var userEntity domain.User
	err := r.db.Preload("Merchants").Where("user_name = ?", username).First(&userEntity).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &userEntity, err
}

func (r userRepository) CreateUser(user domain.User) (domain.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}
