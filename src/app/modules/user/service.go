package user

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input UserInput) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input UserInput) (User, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return User{}, err
	}

	user := User{
		Username:    input.Username,
		Email:       input.Email,
		Password:    string(password),
		Role:        input.Role,
		CreatedDate: time.Now(),
		CreatedBy:   int(uuid.New().ID()),
	}
	err = s.repository.Save(&user)
	if err != nil {
		return user, nil
	}
	return user, nil
}
