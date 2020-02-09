package adding

import (
	"errors"
	"testing"

	"github.com/joaodubas/deliveramate/pkg/storage"
)

func TestAddPartner(t *testing.T) {
	s := NewService(&repository{})

	testCases := []struct {
		name    string
		partner storage.Partner
		err     error
		id      int
	}{
		{
			"Success",
			storage.Partner{
				ID:          1,
				TradingName: "Sample 1",
				OwnerName:   "Sample 1",
				Document:    "11.111.111/1111-00",
			},
			nil,
			1,
		},
		{
			"Fail Duplicate ID",
			storage.Partner{
				ID: 1000,
			},
			storage.ErrorDuplicateID,
			1000,
		},
		{
			"Fail Duplicate Document",
			storage.Partner{
				Document: "00.000.000/0000-00",
			},
			storage.ErrorDuplicateDocument,
			1,
		},
	}

	for _, d := range testCases {
		t.Run(d.name, func(t *testing.T) {
			p, err := s.AddPartner(d.partner)

			if d.err != nil && err == nil {
				t.Errorf("AddPartner should fail for partner id %d", d.partner.ID)
			}

			if d.err != nil && !errors.Is(err, d.err) {
				t.Errorf("AddPartner expected error %v | got error %v", d.err, err)
			}

			if p.ID != d.partner.ID {
				t.Errorf("AddPartner expected partner %d | got partner %d", d.partner.ID, p.ID)
			}
		})
	}
}

type repository struct{}

func (r *repository) AddPartner(p storage.Partner) (storage.Partner, error) {
	if p.ID == 1000 {
		return p, storage.ErrorDuplicateID
	}

	if p.Document == "00.000.000/0000-00" {
		return p, storage.ErrorDuplicateDocument
	}

	return p, nil
}
