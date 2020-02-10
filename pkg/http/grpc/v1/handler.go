package v1

import (
	"context"

	"github.com/joaodubas/deliveramate/pkg/adding"
	"github.com/joaodubas/deliveramate/pkg/listing"
)

type service struct {
	adder  adding.Service
	lister listing.Service
}

func NewService(addService adding.Service, listService listing.Service) *service {
	return &service{addService, listService}
}

func (s *service) CreatePartner(ctx context.Context, r *CreateRequest) (*CreateResponse, error) {
	sp, err := toStoragePartner(r.GetPartner())
	if err != nil {
		return &CreateResponse{Api: r.GetApi(), Partner: r.GetPartner()}, err
	}

	sp, err = s.adder.AddPartner(sp)
	if err != nil {
		return &CreateResponse{Api: r.GetApi(), Partner: r.GetPartner()}, err
	}

	p, err := fromStoragePartner(sp)
	if err != nil {
		return &CreateResponse{Api: r.GetApi(), Partner: p}, err
	}

	return &CreateResponse{Api: r.GetApi(), Partner: p}, nil
}

func (s *service) GetPartner(ctx context.Context, r *GetRequest) (*GetResponse, error) {
	sp, err := s.lister.GetPartnerByID(int(r.GetId()))
	if err != nil {
		return &GetResponse{Api: r.GetApi(), Partner: &Partner{}}, err
	}

	p, err := fromStoragePartner(sp)
	if err != nil {
		return &GetResponse{Api: r.GetApi(), Partner: p}, err
	}

	return &GetResponse{Api: r.GetApi(), Partner: p}, nil
}

func (s *service) FilterPartnerByLocation(ctx context.Context, r *FilterLocationRequest) (*FilterLocationResponse, error) {
	return &FilterLocationResponse{}, nil
}
