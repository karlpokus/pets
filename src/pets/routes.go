package pets

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"log"


	"pets/internal/model"
	"pets/internal/store"
)

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
		fmt.Fprintf(w, "%s", pets)
	}
}

func addPet(s store.Pets) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var pet model.Pet
		if err := json.NewDecoder(r.Body).Decode(&pet); err != nil {
			http.Error(w, http.StatusText(400), 400)
			return
		}
		defer r.Body.Close()
		ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
		defer cancel()
		if err := s.InsertOne(ctx, pet); err != nil {
			http.Error(w, http.StatusText(400), 400)
			return
		}
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
