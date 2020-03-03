package listing

import (
	"github.com/joaodubas/deliveramate/pkg/storage"
	geojson "github.com/paulmach/go.geojson"
)

type Repository interface {
	GetPartnerByID(int) (storage.Partner, error)
	FilterPartnersByLocation(geojson.Geometry) ([]storage.Partner, error)
}

type Service interface {
	GetPartnerByID(int) (storage.Partner, error)
	FilterPartnersByLocation(geojson.Geometry) ([]storage.Partner, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) GetPartnerByID(id int) (storage.Partner, error) {
	return s.repo.GetPartnerByID(id)
}

func (s *service) FilterPartnersByLocation(point geojson.Geometry) ([]storage.Partner, error) {
	return s.repo.FilterPartnersByLocation(point)
}
