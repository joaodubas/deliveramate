package storage

import geojson "github.com/paulmach/go.geojson"

type Partner struct {
	ID           int              `json:"id"`
	TradingName  string           `json:"tradingName"`
	OwnerName    string           `json:"ownerName"`
	Document     string           `json:"document"`
	CoverageArea geojson.Geometry `json:"coverageArea"`
	Address      geojson.Geometry `json:"address"`
}
