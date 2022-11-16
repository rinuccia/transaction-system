package service

import (
	"github.com/rinuccia/transaction-system/internal/models"
	"github.com/rinuccia/transaction-system/internal/repository/postgres"
)

type userService struct {
	repo postgres.UserRepository
}

func newUserService(r postgres.UserRepository) *userService {
	return &userService{
		repo: r,
	}
}

func (u *userService) GetUser(userID string) (models.User, error) {
	user, err := u.repo.FindUser(userID)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (u *userService) CreateUser(user models.User) (string, error) {
	userID, err := u.repo.InsertUser(user)
	if err != nil {
		return "", err
	}
	return userID, nil
}
