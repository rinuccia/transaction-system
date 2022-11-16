package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/rinuccia/transaction-system/internal/models"
)

type UserRepository interface {
	FindUser(userID string) (models.User, error)
	InsertUser(user models.User) (string, error)
	Replenish(request models.UserRequest)
	Withdraw(request models.UserRequest)
}

type Repository struct {
	UserRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		newUserRepo(db),
	}
}
