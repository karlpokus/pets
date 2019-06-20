package pets

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"go.elastic.co/apm/module/apmhttprouter"
	"go.mongodb.org/mongo-driver/mongo"
	"pets/internal/db"
)

type Server struct {
	*http.Server
	*mongo.Client
	Addr string
}

var Version = "vX.Y.Z" // injected at build time

var Stdout = log.New(os.Stdout, "", 0)
var Stderr = log.New(os.Stderr, "", 0)

func cleanupOnExit(s *Server) {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	<-sig
	Stdout.Println()
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	if err := s.Client.Disconnect(ctx); err != nil {
		Stderr.Printf("Client disconnect err: %s", err)
	} else {
		Stdout.Println("db closed")
	}
	if err := s.Server.Shutdown(ctx); err != nil {
		Stderr.Printf("Server shutdown err: %s", err)
	} else {
		Stdout.Println("Server shutdown complete")
	}
}

func (s *Server) Start() error {
	go cleanupOnExit(s)
	Stdout.Println(fmt.Sprintf("pets %s listening on %s", Version, s.Addr))
	return s.ListenAndServe()
}

func New() (*Server, error) {
	client, err := db.New()
	if err != nil {
		return nil, err
	}
	Stdout.Println("connected to db")
	router := apmhttprouter.New() // wraps httprouter
	router.Handler("GET", "/api/v1/pets", logRequest(getPets(client)))
	router.Handler("GET", "/api/v1/stats", stats())
	router.Handler("POST", "/api/v1/pet", logRequest(addPet(client)))
	router.Handler("GET", "/api/v1/ping", ping(client))

	host := os.Getenv("HTTP_HOST")
	port := os.Getenv("HTTP_PORT")
	addr := fmt.Sprintf("%s:%s", host, port)
	return &Server{
		Server: &http.Server{
			Addr:              addr,
			Handler:           router,
			ReadTimeout:       10 * time.Second,
			WriteTimeout:      10 * time.Second,
			ReadHeaderTimeout: 10 * time.Second,
			MaxHeaderBytes:    1 << 20, // 1 MB
		},
		Addr:   addr,
		Client: client,
	}, nil
}
