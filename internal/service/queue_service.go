package service

import (
	"github.com/rinuccia/transaction-system/config"
	"github.com/rinuccia/transaction-system/internal/models"
	"github.com/rinuccia/transaction-system/internal/repository/postgres"
	"sync"
)

type queueService struct {
	mx      sync.Mutex
	wg      sync.WaitGroup
	m       map[string]chan models.UserRequest
	bufSize int
	repo    postgres.UserRepository
}

func newQueueService(r postgres.UserRepository, cfg *config.Config) *queueService {
	return &queueService{
		m:       make(map[string]chan models.UserRequest),
		bufSize: cfg.MaxQueueSize,
		repo:    r,
	}
}

func (s *queueService) handle(ch chan models.UserRequest) {
	defer s.wg.Done()
	for request := range ch {
		if request.OperationCode == "+" {
			s.repo.Replenish(request)
		}
		if request.OperationCode == "-" {
			s.repo.Withdraw(request)
		}
	}
}

func (s *queueService) Handle(req models.UserRequest) {
	if ch := s.getChan(req.UserId); ch != nil {
		ch <- req
	}
}

func (s *queueService) getChan(userID string) chan models.UserRequest {
	s.mx.Lock()
	defer s.mx.Unlock()
	if val, ok := s.m[userID]; ok {
		return val
	}

	if _, err := s.repo.FindUser(userID); err != nil {
		return nil
	}

	ch := make(chan models.UserRequest, s.bufSize)

	s.wg.Add(1)
	go s.handle(ch)

	s.m[userID] = ch

	return ch
}

func (s *queueService) Close() {
	s.mx.Lock()
	defer s.mx.Unlock()
	for _, val := range s.m {
		close(val)
	}
	s.wg.Wait()
}
