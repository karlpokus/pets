package pets

import (
	"fmt"
	"log"
	"context"
	"time"
	"strings"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
)

type Pet struct {
  Name, Kind string
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

func getPets(client *mongo.Client, stderr *log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		coll := client.Database("pets").Collection("pets")
		ctx, cancel := context.WithTimeout(r.Context(), 3 * time.Second)
	  defer cancel()
	  cur, err := coll.Find(ctx, bson.M{}, options.Find())
	  if err != nil {
			stderr.Printf("Unable to find pets: %s", err)
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
	    return
	  }
	  defer cur.Close(ctx)
	  var list []Pet

	  for cur.Next(ctx) { // cur.All(ctx, &list) is undefined. Bug submitted
	    var p Pet
	    if err := cur.Decode(&p); err != nil {
	      stderr.Printf("Pets cursor decode err: %s", err)
				http.Error(w, http.StatusText(500), http.StatusInternalServerError)
	      return
	    }
	    list = append(list, p)
	  }
	  if err := cur.Err(); err != nil {
	    stderr.Printf("Pets cursor err: %s", err) // log and move on
	  }
		// only return names for now
		var names []string
	  for _, p := range list {
	    names = append(names, p.Name)
	  }
		fmt.Fprintf(w, "%s\n", strings.Join(names, ","))
	}
}
