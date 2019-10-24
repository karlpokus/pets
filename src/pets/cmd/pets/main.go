package main

import (
	"flag"

  "pets"
	"github.com/karlpokus/srv"
)

// set to true when running natively
var native = flag.Bool("n", false, "running natively")

func main() {
	flag.Parse()
	stdout := pets.Logging(*native)
  s, err := srv.New(pets.Conf(*native, stdout))
  if err != nil {
    stdout.Fatal(err)
  }
  err = s.Start()
	if err != nil {
		stdout.Fatal(err)
	}
	stdout.Println("main exited")
}
