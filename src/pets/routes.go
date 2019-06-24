package pets

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"runtime"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"pets/internal/model"
)

type Stats struct {
	Name      string `json:"name"`
	Version   string `json:"version"`
	GoVersion string `json:"go_version"`
}

func logRequest(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		Stdout.Print(fmt.Sprintf("%s %s", r.Method, r.URL.Path))
		next.ServeHTTP(w, r)
	}
}

func stats() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		payload := Stats{
			Name:      "pets",
			Version:   Version,
			GoVersion: runtime.Version(),
		}
		var buf bytes.Buffer
		err := json.NewEncoder(&buf).Encode(payload)
		if err != nil {
			Stderr.Printf("Unable to decode stats payload: %s", err)
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(buf.Bytes())
	}
}

func getPets(client model.MongoClient) http.HandlerFunc { // *mongo.Client
	return func(w http.ResponseWriter, r *http.Request) {
		coll := client.Database("pets").Collection("pets")
		ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
		defer cancel()
		cur, err := coll.Find(ctx, bson.M{})
		if err != nil {
			Stderr.Printf("Unable to find pets: %s", err)
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
			return
		}
		defer cur.Close(ctx)
		var list []model.Pet

		for cur.Next(ctx) { // cur.All(ctx, &list) is undefined. Bug submitted
			var p model.Pet
			if err := cur.Decode(&p); err != nil {
				Stderr.Printf("Pets cursor decode err: %s", err)
				http.Error(w, http.StatusText(500), http.StatusInternalServerError)
				return
			}
			list = append(list, p)
		}
		if err := cur.Err(); err != nil {
			Stderr.Printf("Pets cursor err: %s", err) // log and move on
		}
		// only return names for now
		var names []string
		for _, p := range list {
			names = append(names, p.Name)
		}
		fmt.Fprintf(w, "%s\n", strings.Join(names, ","))
	}
}

func addPet(client model.MongoClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var p model.Pet
		if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
			Stderr.Printf("Unable to read body: %s", err)
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		coll := client.Database("pets").Collection("pets")
		ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
		defer cancel()

		err := coll.InsertOne(ctx, p)
		if err != nil {
			Stderr.Printf("Unable to insert pet: %s", err)
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "%s added", p.Name)
	}
}

func ping(client model.MongoClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()
		err := client.Ping(ctx)
		if err != nil {
			Stderr.Printf("db ping failure: %s", err)
			http.Error(w, "db ping failure", http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "ok")
	}
}
