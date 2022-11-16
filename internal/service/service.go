package service

import (
	"github.com/rinuccia/transaction-system/config"
	"github.com/rinuccia/transaction-system/internal/repository/postgres"
)

type Service struct {
	User
	Queue
}

func NewService(repos *postgres.Repository, cfg *config.Config) *Service {
	return &Service{
		newUserService(repos.UserRepository),
		newQueueService(repos.UserRepository, cfg),
	}
}
