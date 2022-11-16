package service

import (
	"github.com/rinuccia/transaction-system/internal/models"
)

type (
	User interface {
		GetUser(userID string) (models.User, error)
		CreateUser(user models.User) (string, error)
	}

	Queue interface {
		Handle(req models.UserRequest)
		Close()
	}
)
