package models

import (
	"backend/internal/entities"

	"gorm.io/gorm"
)

type UserModel interface {
	Create(user *entities.User) error
	GetByEmailOrUsername(identifier string) (*entities.User, error)
}

type userModel struct {
	db *gorm.DB
}

func NewUserModel(db *gorm.DB) UserModel {
	return &userModel{db: db}
}

func (r *userModel) Create(user *entities.User) error {
	return r.db.Create(user).Error
}

func (r *userModel) GetByEmailOrUsername(identifier string) (*entities.User, error) {
	var user entities.User
	err := r.db.Where("email = ? OR user_name = ?", identifier, identifier).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
