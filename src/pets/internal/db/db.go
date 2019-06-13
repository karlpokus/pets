// db returns a connection to a mongodb instance
package db

import (
  "fmt"
  "context"
  "time"
  "os"

  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
  "go.elastic.co/apm/module/apmmongo"
)

// New returns a validated client connection based on options read from env vars
func New() (*mongo.Client, error) {
  clientOpts := options.Client().SetAppName("pets-service")
  clientOpts.SetAuth(options.Credential{
    AuthSource: os.Getenv("MONGODB_PETS_AUTHSOURCE"),
    Username: os.Getenv("MONGODB_PETS_USER"),
    Password: os.Getenv("MONGODB_PETS_PWD"),
  })
  clientOpts.SetMonitor(apmmongo.CommandMonitor())
  clientOpts.ApplyURI(fmt.Sprintf("mongodb://localhost:%s", os.Getenv("MONGODB_PORT_DEV")))

  ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
  defer cancel()
  client, err := mongo.Connect(ctx, clientOpts)
  // connect err cannot be trusted!
  // is nil even though ping ctx deadline exceeded
  // tunnel was not started so there was no mongod running on the specified port
  // TODO: report bug
  err = client.Ping(ctx, nil)
  if err != nil {
    err = fmt.Errorf("mongodb connection err: %s", err)
  }
  return client, err
}
