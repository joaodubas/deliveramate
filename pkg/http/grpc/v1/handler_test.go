package v1

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/joaodubas/deliveramate/pkg/adding"
	"github.com/joaodubas/deliveramate/pkg/listing"
	"github.com/joaodubas/deliveramate/pkg/storage"
	geojson "github.com/paulmach/go.geojson"
)

func TestCreatePartner(t *testing.T) {
	testCases := []struct {
		name    string
		partner Partner
		err     error
	}{
		{
			"success",
			Partner{
				Id:          1,
				TradingName: "Sample 1",
				OwnerName:   "Sample 1",
				Document:    "00.000.000/0000-00",
				CoverageArea: []byte(`{
					"type": "Polygon",
					"coordinates": [[
						[-46.720004081726074, -23.533084428991376],
						[-46.71858787536621, -23.533084428991376],
						[-46.71858787536621, -23.532041754244908],
						[-46.720004081726074, -23.532041754244908],
						[-46.720004081726074, -23.533084428991376]
					]]
				}`),
				Address: []byte(`{
					"type": "Point",
					"coordinates": [-46.71938180923462, -23.53242538082001]
				}`),
			},
			nil,
		},
		{
			"fail wrong coverage",
			Partner{
				Id:          2,
				TradingName: "Sample 2",
				OwnerName:   "Sample 2",
				Document:    "11.111.111/1111-00",
				CoverageArea: []byte(`{
					"type": "Point",
					"coordinates": [-46.71938180923462, -23.53242538082001]
				}`),
				Address: []byte(`{
					"type": "Point",
					"coordinates": [-46.71938180923462, -23.53242538082001]
				}`),
			},
			storage.ErrorWrongCoverageArea,
		},
		{
			"fail wrong address",
			Partner{
				Id:          3,
				TradingName: "Sample 3",
				OwnerName:   "Sample 3",
				Document:    "22.222.222/2222-00",
				CoverageArea: []byte(`{
					"type": "Polygon",
					"coordinates": [[
						[-46.720004081726074, -23.533084428991376],
						[-46.71858787536621, -23.533084428991376],
						[-46.71858787536621, -23.532041754244908],
						[-46.720004081726074, -23.532041754244908],
						[-46.720004081726074, -23.533084428991376]
					]]
				}`),
				Address: []byte(`{
					"type": "LineString",
					"coordinates": [
						[-46.71878635883331, -23.533384442363957],
						[-46.718958020210266, -23.53286802547826]
					]
				}`),
			},
			storage.ErrorWrongAddress,
		},
	}

	for _, d := range testCases {
		t.Run(d.name, func(t *testing.T) {
			p, err := srv.CreatePartner(
				context.TODO(),
				&CreateRequest{Api: "v1", Partner: &d.partner},
			)
			if d.err != nil && err == nil {
				t.Errorf("CreatePartner: should've failed %v", p)
			} else if d.err != nil && !errors.Is(err, d.err) {
				t.Errorf("CreatePartner: expected error %v | got error %v", d.err, err)
			} else if d.err == nil && d.partner.Id != p.Partner.Id {
				t.Errorf("CreatePartner: expected id %d | got id %d", d.partner.Id, p.Partner.Id)
			}
		})
	}
}

func TestGetPartner(t *testing.T) {
	testCases := []struct {
		name string
		id   int64
		err  error
	}{
		{
			"success",
			1,
			nil,
		},
		{
			"fail",
			1000,
			storage.ErrorNotFound,
		},
	}

	for _, d := range testCases {
		t.Run(d.name, func(t *testing.T) {
			p, err := srv.GetPartner(
				context.TODO(),
				&GetRequest{Api: "v1", Id: d.id},
			)
			if d.err != nil && err == nil {
				t.Errorf("GetPartner: should've failed for id %d", d.id)
			} else if d.err != nil && !errors.Is(err, d.err) {
				t.Errorf("GetPartner: expected error %v | got error %v", d.err, err)
			} else if d.err == nil && d.id != p.Partner.GetId() {
				t.Errorf("GetPartner: expected id %d | got id %d", d.id, p.Partner.GetId())
			}
		})
	}
}

func TestFilterPartnerByLocation(t *testing.T) {

}

var srv = NewService(adding.NewService(&repo{}), listing.NewService(&repo{}))

type repo struct{}

func (r *repo) AddPartner(p storage.Partner) (storage.Partner, error) {
	return p, nil
}

func (r *repo) GetPartnerByID(id int) (storage.Partner, error) {
	if id == 1000 {
		return storage.Partner{}, storage.ErrorNotFound
	}
	return storage.Partner{
		ID:          id,
		TradingName: fmt.Sprintf("Sample %d", id),
		OwnerName:   fmt.Sprintf("Sample %d", id),
		Document:    "00.000.000/0000-00",
		CoverageArea: geojson.Geometry{
			Type: geojson.GeometryMultiPolygon,
			MultiPolygon: [][][][]float64{{{
				{-46.720004081726074, -23.533084428991376},
				{-46.71858787536621, -23.533084428991376},
				{-46.71858787536621, -23.532041754244908},
				{-46.720004081726074, -23.532041754244908},
				{-46.720004081726074, -23.533084428991376},
			}}},
		},
		Address: geojson.Geometry{
			Type:  geojson.GeometryPoint,
			Point: []float64{-46.71938180923462, -23.53242538082001},
		},
	}, nil
}

func (r *repo) FilterPartnerByLocation(p geojson.Geometry) ([]storage.Partner, error) {
	return []storage.Partner{}, nil
}
