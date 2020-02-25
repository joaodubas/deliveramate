package tile38

import (
	"errors"
	"testing"

	"github.com/go-redis/redis/v7"
	"github.com/joaodubas/deliveramate/pkg/storage"
	geojson "github.com/paulmach/go.geojson"
)

var s, _ = NewStorage("db_test:9851")

func TestAddPartner(t *testing.T) {
	testCases := []struct {
		name    string
		partner storage.Partner
		err     error
		id      int
	}{
		{
			"AddPartner success",
			storage.Partner{
				ID:          1,
				TradingName: "Sample 1",
				OwnerName:   "Owner 1",
				Document:    "00.000.000/0000-00",
				CoverageArea: geojson.Geometry{
					Type: geojson.GeometryMultiPolygon,
					MultiPolygon: [][][][]float64{{{
						{-46.719199419021606, -23.53602551417083},
						{-46.71830892562866, -23.53492384448112},
						{-46.718287467956536, -23.534628752819078},
						{-46.719253063201904, -23.531874532054413},
						{-46.71980023384094, -23.531500740507177},
						{-46.72041177749634, -23.531412210774658},
						{-46.72041177749634, -23.53323197663715},
						{-46.719199419021606, -23.53602551417083},
					}}},
				},
				Address: geojson.Geometry{
					Type:  geojson.GeometryPoint,
					Point: []float64{-46.71938180923462, -23.53242538082001},
				},
			},
			nil,
			1,
		},
		{
			"AddPartner success: format cnpj document",
			storage.Partner{
				ID:          3,
				TradingName: "Sample 3",
				OwnerName:   "Owner 3",
				Document:    "33333333333300",
				CoverageArea: geojson.Geometry{
					Type: geojson.GeometryMultiPolygon,
					MultiPolygon: [][][][]float64{{{
						{-46.719199419021606, -23.53602551417083},
						{-46.71830892562866, -23.53492384448112},
						{-46.718287467956536, -23.534628752819078},
						{-46.719253063201904, -23.531874532054413},
						{-46.71980023384094, -23.531500740507177},
						{-46.72041177749634, -23.531412210774658},
						{-46.72041177749634, -23.53323197663715},
						{-46.719199419021606, -23.53602551417083},
					}}},
				},
				Address: geojson.Geometry{
					Type:  geojson.GeometryPoint,
					Point: []float64{-46.71938180923462, -23.53242538082001},
				},
			},
			nil,
			3,
		},
		{
			"AddPartner success: format cpf document",
			storage.Partner{
				ID:          4,
				TradingName: "Sample 4",
				OwnerName:   "Owner 4",
				Document:    "44444444400",
				CoverageArea: geojson.Geometry{
					Type: geojson.GeometryMultiPolygon,
					MultiPolygon: [][][][]float64{{{
						{-46.719199419021606, -23.53602551417083},
						{-46.71830892562866, -23.53492384448112},
						{-46.718287467956536, -23.534628752819078},
						{-46.719253063201904, -23.531874532054413},
						{-46.71980023384094, -23.531500740507177},
						{-46.72041177749634, -23.531412210774658},
						{-46.72041177749634, -23.53323197663715},
						{-46.719199419021606, -23.53602551417083},
					}}},
				},
				Address: geojson.Geometry{
					Type:  geojson.GeometryPoint,
					Point: []float64{-46.71938180923462, -23.53242538082001},
				},
			},
			nil,
			4,
		},
		{
			"AddPartner failure: duplicate id",
			storage.Partner{
				ID:          1,
				TradingName: "Repeat 1",
				OwnerName:   "Repeat 1",
				Document:    "11.111.111/1111-00",
				CoverageArea: geojson.Geometry{
					Type: geojson.GeometryMultiPolygon,
					MultiPolygon: [][][][]float64{{{
						{-46.719199419021606, -23.53602551417083},
						{-46.71830892562866, -23.53492384448112},
						{-46.718287467956536, -23.534628752819078},
						{-46.719253063201904, -23.531874532054413},
						{-46.71980023384094, -23.531500740507177},
						{-46.72041177749634, -23.531412210774658},
						{-46.72041177749634, -23.53323197663715},
						{-46.719199419021606, -23.53602551417083},
					}}},
				},
				Address: geojson.Geometry{
					Type:  geojson.GeometryPoint,
					Point: []float64{-46.71938180923462, -23.53242538082001},
				},
			},
			storage.ErrorDuplicateID,
			0,
		},
		{
			"AddPartner failure: duplicate document",
			storage.Partner{
				ID:          2,
				TradingName: "Repeat 2",
				OwnerName:   "Repeat 2",
				Document:    "00.000.000/0000-00",
				CoverageArea: geojson.Geometry{
					Type: geojson.GeometryMultiPolygon,
					MultiPolygon: [][][][]float64{{{
						{-46.719199419021606, -23.53602551417083},
						{-46.71830892562866, -23.53492384448112},
						{-46.718287467956536, -23.534628752819078},
						{-46.719253063201904, -23.531874532054413},
						{-46.71980023384094, -23.531500740507177},
						{-46.72041177749634, -23.531412210774658},
						{-46.72041177749634, -23.53323197663715},
						{-46.719199419021606, -23.53602551417083},
					}}},
				},
				Address: geojson.Geometry{
					Type:  geojson.GeometryPoint,
					Point: []float64{-46.71938180923462, -23.53242538082001},
				},
			},
			storage.ErrorDuplicateDocument,
			0,
		},
	}

	for _, d := range testCases {
		t.Run(d.name, func(t *testing.T) {
			p, err := s.AddPartner(d.partner)
			if d.err != nil && err == nil {
				t.Errorf("AddPartner must fail for id (%d)", d.partner.ID)
			}
			if d.err != nil && !errors.Is(err, d.err) {
				t.Errorf("AddPartner expected error (%s) got error (%s)", d.err, err)
			}
			if d.id > 0 && d.id != p.ID {
				t.Errorf("AddPartner should persist id (%d)", d.partner.ID)
			}
		})
	}

	drop(s)
}

func TestGetPartnerByID(t *testing.T) {
	partners := []storage.Partner{
		storage.Partner{
			ID:          10,
			TradingName: "Sample 10",
			OwnerName:   "Owner 10",
			Document:    "00.000.000/0000-11",
			CoverageArea: geojson.Geometry{
				Type: geojson.GeometryMultiPolygon,
				MultiPolygon: [][][][]float64{{{
					{-46.719199419021606, -23.53602551417083},
					{-46.71830892562866, -23.53492384448112},
					{-46.718287467956536, -23.534628752819078},
					{-46.719253063201904, -23.531874532054413},
					{-46.71980023384094, -23.531500740507177},
					{-46.72041177749634, -23.531412210774658},
					{-46.72041177749634, -23.53323197663715},
					{-46.719199419021606, -23.53602551417083},
				}}},
			},
			Address: geojson.Geometry{
				Type:  geojson.GeometryPoint,
				Point: []float64{-46.71938180923462, -23.53242538082001},
			},
		},
		storage.Partner{
			ID:          20,
			TradingName: "Sample 20",
			OwnerName:   "Owner 20",
			Document:    "11.111.111/1111-11",
			CoverageArea: geojson.Geometry{
				Type: geojson.GeometryMultiPolygon,
				MultiPolygon: [][][][]float64{{{
					{-46.719199419021606, -23.53602551417083},
					{-46.71830892562866, -23.53492384448112},
					{-46.718287467956536, -23.534628752819078},
					{-46.719253063201904, -23.531874532054413},
					{-46.71980023384094, -23.531500740507177},
					{-46.72041177749634, -23.531412210774658},
					{-46.72041177749634, -23.53323197663715},
					{-46.719199419021606, -23.53602551417083},
				}}},
			},
			Address: geojson.Geometry{
				Type:  geojson.GeometryPoint,
				Point: []float64{-46.71938180923462, -23.53242538082001},
			},
		},
	}

	for _, p := range partners {
		if p, err := s.AddPartner(p); err != nil {
			t.Errorf("GetPartnerByID: fail to add expected partner %v", p)
		}
	}

	testCases := []struct {
		name string
		id   int
		err  error
	}{
		{
			"success id 10",
			10,
			nil,
		},
		{
			"success id 20",
			20,
			nil,
		},
		{
			"not found",
			30,
			storage.ErrorNotFound,
		},
	}

	for _, d := range testCases {
		t.Run(d.name, func(t *testing.T) {
			p, err := s.GetPartnerByID(d.id)
			if err != nil && d.err == nil {
				t.Errorf("GetPartnerByID: should not fail %v", err)
			} else if err != nil && !errors.Is(err, d.err) {
				t.Errorf("GetPartnerByID: expected error %v | got error %v", d.err, err)
			} else if d.err == nil && p.ID != d.id {
				t.Errorf("GetPartnerByID: ID expected %d | got %d", d.id, p.ID)
			}
		})
	}

	drop(s)
}

func TestFilterPartnersByLocation(t *testing.T) {
	partners := []storage.Partner{
		storage.Partner{
			ID:          100,
			TradingName: "Sample 100",
			OwnerName:   "Owner 100",
			Document:    "00.000.000/0000-22",
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
		},
		storage.Partner{
			ID:          200,
			TradingName: "Sample 200",
			OwnerName:   "Owner 200",
			Document:    "11.111.111/1111-22",
			CoverageArea: geojson.Geometry{
				Type: geojson.GeometryMultiPolygon,
				MultiPolygon: [][][][]float64{{{
					{-46.720519065856934, -23.532592602310373},
					{-46.719017028808594, -23.532592602310373},
					{-46.719017028808594, -23.53145155732979},
					{-46.720519065856934, -23.53145155732979},
					{-46.720519065856934, -23.532592602310373},
				}}},
			},
			Address: geojson.Geometry{
				Type:  geojson.GeometryPoint,
				Point: []float64{-46.71938180923462, -23.53242538082001},
			},
		},
		storage.Partner{
			ID:          300,
			TradingName: "Sample 300",
			OwnerName:   "Owner 300",
			Document:    "22.222.222/2222-22",
			CoverageArea: geojson.Geometry{
				Type: geojson.GeometryMultiPolygon,
				MultiPolygon: [][][][]float64{{{
					{-46.71961784362792, -23.532464727072163},
					{-46.71857714653015, -23.532464727072163},
					{-46.71857714653015, -23.531431884053692},
					{-46.71961784362792, -23.531431884053692},
					{-46.71961784362792, -23.532464727072163},
				}}},
			},
			Address: geojson.Geometry{
				Type:  geojson.GeometryPoint,
				Point: []float64{-46.71938180923462, -23.53242538082001},
			},
		},
		storage.Partner{
			ID:          400,
			TradingName: "Sample 400",
			OwnerName:   "Owner 400",
			Document:    "33.333.333/3333-22",
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
	}

	for _, p := range partners {
		if _, err := s.AddPartner(p); err != nil {
			t.Errorf("FilterPartnersByLocation: fail to add expected partner %v", p)
		}
	}

	testCases := []struct {
		name     string
		location geojson.Geometry
		expected int
		err      error
	}{
		{
			"fetch 3 partners",
			*geojson.NewPointGeometry([]float64{-46.71938180923462, -23.53242538082001}),
			3,
			nil,
		},
		{
			"fetch 1 partner",
			*geojson.NewPointGeometry([]float64{-46.719521284103394, -23.530999071235296}),
			1,
			nil,
		},
		{
			"fetch 0 partner",
			*geojson.NewPointGeometry([]float64{-46.718266010284424, -23.531333517629083}),
			0,
			nil,
		},
		{
			"fail: wrong geometry",
			*geojson.NewPolygonGeometry([][][]float64{{
				{-46.718266010284424, -23.531333517629083},
				{-46.718266010284424, -23.531333517629083},
			}}),
			0,
			storage.ErrorWrongAddress,
		},
	}
	for _, d := range testCases {
		t.Run(d.name, func(t *testing.T) {
			ps, err := s.FilterPartnersByLocation(d.location)
			if d.err != nil && !errors.Is(err, d.err) {
				t.Errorf("FilterPartnersByLocation: expected error %v | got error %v", d.err, err)
			}
			if d.err == nil && err != nil {
				t.Errorf("FilterPartnersByLocation: should not fail (%v)", err)
			}
			if len(ps) != d.expected {
				t.Errorf("FilterParnersByLocation: expected %d | got %d", d.expected, len(ps))
			}
		})
	}

	drop(s)
}

func drop(s *Storage) {
	drop := func(key string) error {
		cmd := redis.NewStringCmd("DROP", key)
		_ = s.db.Process(cmd)
		_, err := cmd.Result()
		return err
	}

	_ = drop("partner:id:document")
	_ = drop("partner:document:id")
	_ = drop("partner")
	_ = drop("partner:coverage")
	_ = drop("partner:address")
}
