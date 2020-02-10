package v1

import (
	"encoding/json"

	"github.com/joaodubas/deliveramate/pkg/storage"
)

func toStoragePartner(p *Partner) (storage.Partner, error) {
	return storage.NewPartner(
		int(p.GetId()),
		p.GetTradingName(),
		p.GetOwnerName(),
		p.GetDocument(),
		p.GetCoverageArea(),
		p.GetAddress(),
	)
}

func fromStoragePartner(sp storage.Partner) (*Partner, error) {
	p := &Partner{}

	coverage, err := json.Marshal(sp.CoverageArea)
	if err != nil {
		return p, err
	}

	address, err := json.Marshal(sp.Address)
	if err != nil {
		return p, err
	}

	p.Id = int64(sp.ID)
	p.TradingName = sp.TradingName
	p.OwnerName = sp.OwnerName
	p.Document = sp.Document
	p.CoverageArea = coverage
	p.Address = address

	return p, nil
}
