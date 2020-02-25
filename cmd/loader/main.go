package main

import (
	"encoding/json"
	"log"

	"github.com/joaodubas/deliveramate/pkg/adding"
	"github.com/joaodubas/deliveramate/pkg/storage"
	"github.com/joaodubas/deliveramate/pkg/storage/tile38"
)

func main() {
	failure := false
	svc := service()
	for i, p := range partners() {
		if np, err := svc.AddPartner(p); err != nil {
			log.Printf("loader: register %d | failure to add partner %d (%s) | error %v", i, np.ID, np.Document, err)
			failure = true
		} else {
			log.Printf("loader: register %d | loaded partner %d (%s)", i, np.ID, np.Document)
		}
	}
	if failure {
		log.Fatalf("loader: failed to load data")
	}
	log.Println("loader: loaded all registers")
}

func service() adding.Service {
	repo, err := tile38.NewStorage("db:9851")
	if err != nil {
		log.Fatalf("loader: error configuring storage (%v)", err)
	}
	return adding.NewService(repo)
}

func partners() []storage.Partner {
	var d []storage.Partner
	if err := json.Unmarshal(data, &d); err != nil {
		log.Fatalf("loader: error converting data (%v)", err)
	}
	return d
}
