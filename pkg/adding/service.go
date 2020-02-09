package adding

import "github.com/joaodubas/deliveramate/pkg/storage"

type Repository interface {
	AddPartner(storage.Partner) (storage.Partner, error)
}

type Service interface {
	AddPartner(storage.Partner) (storage.Partner, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) AddPartner(p storage.Partner) (storage.Partner, error) {
	return s.repo.AddPartner(p)
}
