package pets

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"log"

	"go.elastic.co/apm/module/apmhttprouter"
	"pets/internal/model"
	"pets/internal/store"
)

func routes(s store.Pets, stdout *log.Logger) http.Handler {
	router := apmhttprouter.New() // wraps httprouter
	router.Handler("GET", "/api/v1/pets", logRequest(stdout, getPets(s)))
	router.Handler("POST", "/api/v1/pet", logRequest(stdout, addPet(s)))
	router.Handler("GET", "/api/v1/ping", ping(s))
	return router
}

func logRequest(stdout *log.Logger, next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		stdout.Print(fmt.Sprintf("%s %s", r.Method, r.URL.Path))
		next.ServeHTTP(w, r)
	}
}

func getPets(s store.Pets) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
		defer cancel()
		pets, err := s.Find(ctx)
		if err != nil {
			// TODO: log err
			http.Error(w, http.StatusText(404), 404)
			return
		}
		fmt.Fprintf(w, "%s", pets) // sets Content-Type: application/json; charset=utf-8
	}
}

func addPet(s store.Pets) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var pet model.Pet
		if err := json.NewDecoder(r.Body).Decode(&pet); err != nil {
			http.Error(w, "Malformed json body", 400)
			return
		}
		defer r.Body.Close()
		if !pet.Valid() {
			http.Error(w, "Invalid json body", 400)
			return
		}
		ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
		defer cancel()
		if err := s.InsertOne(ctx, pet); err != nil {
			http.Error(w, http.StatusText(400), 400)
			return
		}
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "%s added", pet.Name)
	}
}

func ping(s store.Pets) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()
		err := s.Ping(ctx)
		if err != nil {
			http.Error(w, "ping failed", 500)
			return
		}
		fmt.Fprintf(w, "pong")
	}
}
