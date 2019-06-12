package main

import (
  "flag"
  "fmt"

  "pets"
)

var (
  version = flag.Bool("version", false, "print version and exit")
  port = flag.String("port", "37042", "http port")
)

func main() {
  flag.Parse()
  if *version {
    fmt.Println(pets.Version)
    return
  }
  server, err := pets.New(*port)
  if err != nil {
    fmt.Printf("Unable to create server: %s", err)
    return
  }
  server.Start()
}
