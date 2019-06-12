// db returns a connection to a mongodb instance
package db

import (
  "fmt"
  "context"
  "time"

  "pets/internal/vault"
  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
  "go.elastic.co/apm/module/apmmongo"
)

type Vault struct {
  Mongodb struct {
    Pets struct {
      User, Pwd, Authsource string
    }
    Port struct {
      Dev string
    }
  }
}

// New returns a client connection based on options read from vault and an error
func New() (client *mongo.Client, err error) {
  var v Vault
  err = vault.View("../../deploy", "group_vars/all", &v)
  if err != nil {
    return
  }
  clientOpts := options.Client().SetAppName("pets-service")
  clientOpts.SetAuth(options.Credential{
    AuthSource: v.Mongodb.Pets.Authsource,
    Username: v.Mongodb.Pets.User,
    Password: v.Mongodb.Pets.Pwd,
  })
  clientOpts.SetMonitor(apmmongo.CommandMonitor())
  clientOpts.ApplyURI(fmt.Sprintf("mongodb://localhost:%s", v.Mongodb.Port.Dev))

  ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
  defer cancel()
  client, err = mongo.Connect(ctx, clientOpts)
  // connect err cannot be trusted!
  // is nil even though ping ctx deadline exceeded
  // tunnel was not started so there was no mongod running on the specified port
  // TODO: report bug
  err = client.Ping(ctx, nil)
  if err != nil {
    err = fmt.Errorf("mongodb connection err: %s", err)
  }
  return
}
