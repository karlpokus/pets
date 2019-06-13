package pets

import (
	"fmt"
	"context"
	"time"
	"strings"
	"net/http"
	"encoding/json"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Pet struct {
  Name string `json:"name"`
	Kind string `json:"kind"`
}

func logRequest(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		Stdout.Print(fmt.Sprintf("%s %s", r.Method, r.URL.Path))
		next.ServeHTTP(w, r)
	}
}

func getVersion() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s\n", Version)
	}
}

func getPets(client *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		coll := client.Database("pets").Collection("pets")
		ctx, cancel := context.WithTimeout(r.Context(), 3 * time.Second)
	  defer cancel()
	  cur, err := coll.Find(ctx, bson.M{}, options.Find())
	  if err != nil {
			Stderr.Printf("Unable to find pets: %s", err)
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
	    return
	  }
	  defer cur.Close(ctx)
	  var list []Pet

	  for cur.Next(ctx) { // cur.All(ctx, &list) is undefined. Bug submitted
	    var p Pet
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

func addPet(client *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		var p Pet
		if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
			Stderr.Printf("Unable to read body: %s", err)
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
	    return
		}
		defer r.Body.Close()

		coll := client.Database("pets").Collection("pets")
		ctx, cancel := context.WithTimeout(r.Context(), 3 * time.Second)
	  defer cancel()

	  res, err := coll.InsertOne(ctx, p, options.InsertOne())
		if err != nil {
			Stderr.Printf("Unable to insert pet: %s", err)
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
	    return
		}
		fmt.Fprintf(w, "new pet added id: %s", res.InsertedID.(primitive.ObjectID).Hex())
	}
}
