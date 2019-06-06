package main

import (
  "flag"
  "os"
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
    os.Exit(0)
  }
  server := pets.New(*port)
  server.Start()
}
