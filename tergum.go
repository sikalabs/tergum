package main

import (
	"errors"
	"flag"
	"log"

	"github.com/sikalabs/tergum/tergum1"
)

func main() {
	// Backup parameters from config file
	path := flag.String("config", "", "tergum config file (json)")

	flag.Parse()

	if *path == "" {
		log.Fatal(errors.New("tergum require config file (-config)"))
	}

	tergum1.Tergum1(*path)
}
