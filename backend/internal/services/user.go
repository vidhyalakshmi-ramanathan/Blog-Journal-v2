package services

import (
	"backend/internal/entities"
	"backend/internal/models"
	"backend/token"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Create(user *entities.User) error
	Login(identifier, password string) (string, error)
}

type userService struct {
	user models.UserModel
}

func NewUserService(user models.UserModel) UserService {
	return &userService{user: user}
}

func (s *userService) Create(user *entities.User) error {
	if user.UserName == "" || user.Email == "" || user.PhoneNumber == "" || user.Password == "" {
		return errors.New("name or email or phonenumber or password are cannot be empty")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return s.user.Create(user)
}

func (s *userService) Login(identifier, password string) (string, error) {
	user, err := s.user.GetByEmailOrUsername(identifier)
	if err != nil {
		return "", errors.New("invalid email/username")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid password")
	}
	token, err := token.CreateToken(user.ID, user.UserName)
	if err != nil {
		return "", errors.New("unauthorized access")
	}
	return token, nil
}
