package main

import (
  "fmt"

  "pets"
)

func main() {
  server, err := pets.New()
  if err != nil {
    fmt.Printf("Unable to create server: %s", err)
    return
  }
  server.Start()
}
