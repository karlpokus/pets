package pets

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"pets/internal/model"
)

type Stats struct {
	Name      string `json:"name"`
	GoVersion string `json:"go_version"`
}

func logRequest(stdout *log.Logger, next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		stdout.Print(fmt.Sprintf("%s %s", r.Method, r.URL.Path))
		next.ServeHTTP(w, r)
	}
}

func getPets(client model.MongoClient) http.HandlerFunc { // *mongo.Client
	return func(w http.ResponseWriter, r *http.Request) {
		coll := client.Database("pets").Collection("pets")
		ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
		defer cancel()
		cur, err := coll.Find(ctx, bson.M{})
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}
		defer cur.Close(ctx)
		var list []model.Pet

		for cur.Next(ctx) { // cur.All(ctx, &list) is undefined. Bug submitted
			var p model.Pet
			if err := cur.Decode(&p); err != nil {
				http.Error(w, http.StatusText(500), 500)
				return
			}
			list = append(list, p)
		}
		if err := cur.Err(); err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
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
			http.Error(w, http.StatusText(400), 400)
			return
		}
		defer r.Body.Close()
		coll := client.Database("pets").Collection("pets")
		ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
		defer cancel()
		err := coll.InsertOne(ctx, p)
		if err != nil {
			http.Error(w, http.StatusText(400), 400)
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
			http.Error(w, "db ping failure", 500)
			return
		}
		fmt.Fprintf(w, "ok")
	}
}
