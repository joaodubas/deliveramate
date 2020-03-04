package main

import (
	"log"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"
  "google.golang.org/grpc/reflection"

	"github.com/joaodubas/deliveramate/pkg/adding"
	v1 "github.com/joaodubas/deliveramate/pkg/http/grpc/v1"
	"github.com/joaodubas/deliveramate/pkg/listing"
	"github.com/joaodubas/deliveramate/pkg/storage/tile38"
)

func main() {
	listen, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("server: error listening to 0.0.0.0:9090: %v", err)
	}

	store, err := tile38.NewStorage(("db:9851"))
	if err != nil {
		log.Fatalf("server: error connecting to storage: %v", err)
	}

	server := grpc.NewServer()
	v1.RegisterPartnerServiceServer(
		server,
		v1.NewService(
			adding.NewService(store),
			listing.NewService(store),
		),
	)
  reflection.Register(server)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			log.Println("server: shutting down grpc")
			server.GracefulStop()
		}
	}()

	log.Println("server: start grpc")
	if err = server.Serve(listen); err != nil {
		log.Fatalf("server: error starting grpc: %v", err)
	}
}
