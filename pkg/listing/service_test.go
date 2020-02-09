package listing

import (
	"errors"
	"fmt"
	"testing"

	"github.com/joaodubas/deliveramate/pkg/storage"
	geojson "github.com/paulmach/go.geojson"
)

func TestGetPartnerByID(t *testing.T) {
	s := NewService(&repository{})

	testCases := []struct {
		name string
		id   int
		err  error
	}{
		{
			"success",
			10,
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
			p, err := s.GetPartnerByID(d.id)
			if d.err != nil && err == nil {
				t.Errorf("GetPartnerByID: should've failed for id %d", d.id)
			} else if d.err != nil && !errors.Is(err, d.err) {
				t.Errorf("GetPartnerById: expected error %w | got error %v", d.err, err)
			} else if d.err == nil && d.id != p.ID {
				t.Errorf("GetPartnerByID: expected id %d | got id %d", d.id, p.ID)
			}
		})
	}
}

func TestFilterParnerByLocation(t *testing.T) {
	s := NewService(&repository{})

	testCases := []struct {
		name   string
		point  geojson.Geometry
		length int
		err    error
	}{
		{
			"success",
			*geojson.NewPointGeometry([]float64{1.0, 1.0}),
			2,
			nil,
		},
		{
			"fail",
			*geojson.NewLineStringGeometry([][]float64{{1.0, 1.0}, {1.0, 1.0}}),
			0,
			storage.ErrorWrongAddress,
		},
	}

	for _, d := range testCases {
		t.Run(d.name, func(t *testing.T) {
			ps, err := s.FilterPartnerByLocation(d.point)
			if d.err != nil && err == nil {
				t.Errorf("FilterPartnerByLocation: should've failed for point %v", d.point)
			} else if d.err != nil && !errors.Is(err, d.err) {
				t.Errorf("FilterPartnerByLocation: expected error %v | got error %v", d.err, err)
			} else if d.err == nil && len(ps) != d.length {
				t.Errorf("FilterPartnerByLocation: expected %d partners | got %d partner", d.length, 2)
			}
		})
	}
}

type repository struct{}

func (r *repository) GetPartnerByID(id int) (storage.Partner, error) {
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
				{-46.72017574310303, -23.531776165960732},
				{-46.71886682510376, -23.531776165960732},
				{-46.71886682510376, -23.530644950597658},
				{-46.72017574310303, -23.530644950597658},
				{-46.72017574310303, -23.531776165960732},
			}}},
		},
		Address: geojson.Geometry{
			Type:  geojson.GeometryPoint,
			Point: []float64{-46.719521284103394, -23.530999071235296},
		},
	}, nil
}

func (r *repository) FilterPartnerByLocation(point geojson.Geometry) ([]storage.Partner, error) {
	ps := []storage.Partner{}

	if !point.IsPoint() {
		return ps, storage.ErrorWrongAddress
	}

	ps = append(
		ps,
		storage.Partner{
			ID:          10,
			TradingName: "Sample 10",
			OwnerName:   "Sample 10",
			Document:    "00.000.000/0000-00",
			CoverageArea: geojson.Geometry{
				Type: geojson.GeometryMultiPolygon,
				MultiPolygon: [][][][]float64{{{
					{-46.72017574310303, -23.531776165960732},
					{-46.71886682510376, -23.531776165960732},
					{-46.71886682510376, -23.530644950597658},
					{-46.72017574310303, -23.530644950597658},
					{-46.72017574310303, -23.531776165960732},
				}}},
			},
			Address: geojson.Geometry{
				Type:  geojson.GeometryPoint,
				Point: []float64{-46.719521284103394, -23.530999071235296},
			},
		},
		storage.Partner{
			ID:          20,
			TradingName: "Sample 20",
			OwnerName:   "Sample 20",
			Document:    "11.111.111/1111-00",
			CoverageArea: geojson.Geometry{
				Type: geojson.GeometryMultiPolygon,
				MultiPolygon: [][][][]float64{{{
					{-46.72017574310303, -23.531776165960732},
					{-46.71886682510376, -23.531776165960732},
					{-46.71886682510376, -23.530644950597658},
					{-46.72017574310303, -23.530644950597658},
					{-46.72017574310303, -23.531776165960732},
				}}},
			},
			Address: geojson.Geometry{
				Type:  geojson.GeometryPoint,
				Point: []float64{-46.719521284103394, -23.530999071235296},
			},
		},
	)

	return ps, nil
}
