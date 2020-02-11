package tile38

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/go-redis/redis/v7"
	"github.com/joaodubas/deliveramate/pkg/storage"
	geojson "github.com/paulmach/go.geojson"
)

type Storage struct {
	db *redis.Client
}

func NewStorage() (*Storage, error) {
	cli := redis.NewClient(&redis.Options{Addr: "db:9851"})
	if _, err := cli.Ping().Result(); err != nil {
		return nil, fmt.Errorf("NewStorage: connection error: %w", err)
	}
	s := new(Storage)
	s.db = cli
	return s, nil
}

func (s *Storage) AddPartner(p storage.Partner) (storage.Partner, error) {
	doc, err := storage.DocumentFormatter(p.Document)
	if err != nil {
		return p, fmt.Errorf("AddPartner: error invalid document (%w)", err)
	}
	p.Document = doc

	if s.existingPartnerID(p.ID) {
		return p, fmt.Errorf("AddPartner: error id already saved (%w)", storage.ErrorDuplicateID)
	}

	if s.existingPartnerDocument(p.Document) {
		return p, fmt.Errorf("AddPartner: error document already saved (%w)", storage.ErrorDuplicateDocument)
	}

	if _, err := s.set(p); err != nil {
		return p, fmt.Errorf("AddPartner: error saving partner (%w)", err)
	}

	return p, nil
}

func (s *Storage) GetPartnerByID(id int) (storage.Partner, error) {
	jgetID, _ := commander(s.db, "JGET", "partner:id:document", id)
	r, err := jgetID.Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return storage.Partner{}, fmt.Errorf("GetPartnerByID: error fetching partner (%w)", storage.ErrorNotFound)
		}
		return storage.Partner{}, fmt.Errorf("GetPartnerID: error fetching ID (%w)", err)
	}

	var d map[string]string
	_ = json.Unmarshal([]byte(r), &d)

	return s.getPartnerByDocument(d["document"])
}

func (s *Storage) FilterPartnersByLocation(point geojson.Geometry) ([]storage.Partner, error) {
	var ps []storage.Partner
	if !point.IsPoint() {
		return ps, fmt.Errorf("FilterPartnersByLocation: wrong type for point (%w)", storage.ErrorWrongAddress)
	}

	p, err := point.MarshalJSON()
	if err != nil {
		return ps, fmt.Errorf("FilterPartnersByLocation: fail to convert point (%w)", err)
	}

	filterID := redis.NewSliceCmd("INTERSECTS", "partner:coverage", "IDS", "OBJECT", string(p))
	_ = s.db.Process(filterID)

	r, err := filterID.Result()
	if err != nil {
		return ps, fmt.Errorf("FilterPartnersByLocation: error fetching identities (%w)", err)
	}

	d := r[1]
	for _, e := range d.([]interface{}) {
		if p, err := s.getPartnerByDocument(e.(string)); err != nil {
			return ps, fmt.Errorf("FilterPartnerByLocation: error getting partner (%w)", err)
		} else {
			ps = append(ps, p)
		}
	}
	return ps, nil
}

func (s *Storage) existingPartnerID(id int) bool {
	jgetID, _ := commander(s.db, "JGET", "partner:id:document", id)
	if r, err := jgetID.Result(); err != nil {
		return false
	} else {
		var d map[string]string
		_ = json.Unmarshal([]byte(r), &d)
		return len(d) > 0
	}
}

func (s *Storage) existingPartnerDocument(document string) bool {
	jgetDocument := redis.NewStringCmd("JGET", "partner:document:id", document)
	_ = s.db.Process(jgetDocument)
	if r, err := jgetDocument.Result(); err != nil {
		return false
	} else {
		var d map[string]int
		_ = json.Unmarshal([]byte(r), &d)
		return len(d) > 0
	}
}

func (s *Storage) getPartnerByDocument(doc string) (storage.Partner, error) {
	jgetPartner, _ := commander(s.db, "JGET", "partner", doc)
	getAddress, _ := commander(s.db, "GET", "partner:address", doc)
	getCoverage, _ := commander(s.db, "GET", "partner:coverage", doc)

	p := storage.Partner{}
	r, err := jgetPartner.Result()
	if err != nil {
		return p, err
	}

	_ = json.Unmarshal([]byte(r), &p)

	address, err := getAddress.Result()
	if err != nil {
		return p, err
	}
	if err = json.Unmarshal([]byte(address), &p.Address); err != nil {
		return p, err
	}

	coverage, err := getCoverage.Result()
	if err != nil {
		return p, err
	}
	if err = json.Unmarshal([]byte(coverage), &p.CoverageArea); err != nil {
		return p, err
	}

	return p, nil
}

func (s *Storage) set(p storage.Partner) (storage.Partner, error) {
	jsetID, _ := commander(s.db, "JSET", "partner:id:document", p.ID, "document", p.Document)
	jsetDocument, _ := commander(s.db, "JSET", "partner:document:id", p.Document, "id", p.ID)
	jsetPartnerID, _ := commander(s.db, "JSET", "partner", p.Document, "id", p.ID)
	jsetPartnerDocument, _ := commander(s.db, "JSET", "partner", p.Document, "document", p.Document)
	jsetPartnerTradingName, _ := commander(s.db, "JSET", "partner", p.Document, "tradingName", p.TradingName)
	jsetPartnerOwnerName, _ := commander(s.db, "JSET", "partner", p.Document, "ownerName", p.OwnerName)

	address, err := json.Marshal(p.Address)
	if err != nil {
		return p, err
	}
	setPoint, _ := commander(s.db, "SET", "partner:address", p.Document, "OBJECT", string(address))
	_ = s.db.Process(setPoint)

	coverage, err := json.Marshal(p.CoverageArea)
	if err != nil {
		return p, err
	}
	setMultiPolygon, _ := commander(s.db, "SET", "partner:coverage", p.Document, "OBJECT", string(coverage))

	if _, err := jsetID.Result(); err != nil {
		return p, fmt.Errorf("set: error setting id (%w)", err)
	} else if _, err := jsetDocument.Result(); err != nil {
		return p, fmt.Errorf("set: error setting document (%w)", err)
	} else if _, err := jsetPartnerID.Result(); err != nil {
		return p, fmt.Errorf("set: error saving details (%w)", err)
	} else if _, err := jsetPartnerDocument.Result(); err != nil {
		return p, fmt.Errorf("set: error saving details (%w)", err)
	} else if _, err := jsetPartnerTradingName.Result(); err != nil {
		return p, fmt.Errorf("set: error saving details (%w)", err)
	} else if _, err := jsetPartnerOwnerName.Result(); err != nil {
		return p, fmt.Errorf("set: error saving details (%w)", err)
	} else if _, err := setPoint.Result(); err != nil {
		return p, fmt.Errorf("set: error saving address (%w)", err)
	} else if _, err := setMultiPolygon.Result(); err != nil {
		return p, fmt.Errorf("set: error saving coverage area (%w)", err)
	}

	return p, nil
}

func commander(client *redis.Client, cmd string, key string, identity interface{}, args ...interface{}) (*redis.StringCmd, error) {
	a := []interface{}{cmd, key, identity}
	a = append(a, args...)
	c := redis.NewStringCmd(a...)
	if err := client.Process(c); err != nil {
		return c, err
	}
	return c, nil
}
