package storage

import (
	"fmt"

	geojson "github.com/paulmach/go.geojson"
)

type Partner struct {
	ID           int              `json:"id"`
	TradingName  string           `json:"tradingName"`
	OwnerName    string           `json:"ownerName"`
	Document     string           `json:"document"`
	CoverageArea geojson.Geometry `json:"coverageArea"`
	Address      geojson.Geometry `json:"address"`
}

func NewPartner(id int, tradingName, ownerName, document string, coverageArea, address []byte) (Partner, error) {
	p := Partner{
		ID:          id,
		TradingName: tradingName,
		OwnerName:   ownerName,
	}

	doc, err := DocumentFormatter(document)
	if err != nil {
		return p, fmt.Errorf("NewPartner: error formatting document (%w)", err)
	}

	coverageGeom, err := geojson.UnmarshalGeometry(coverageArea)
	if err != nil {
		return p, fmt.Errorf("NewPartner: error converting coverage area (%w)", err)
	} else if !coverageGeom.IsMultiPolygon() && !coverageGeom.IsPolygon() {
		return p, fmt.Errorf("NewPartner: invalid coverage area (%w)", ErrorWrongCoverageArea)
	}

	addressGeom, err := geojson.UnmarshalGeometry(address)
	if err != nil {
		return Partner{}, fmt.Errorf("NewPartner: error coverting address (%w)", err)
	} else if !addressGeom.IsPoint() {
		return p, fmt.Errorf("NewPartner: invalid address (%w)", ErrorWrongAddress)
	}

	p.Document = doc
	p.CoverageArea = *coverageGeom
	p.Address = *addressGeom
	return p, nil
}
