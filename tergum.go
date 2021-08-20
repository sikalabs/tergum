package main

import (
	"errors"
	"flag"
	"log"

	"github.com/sikalabs/tergum/tergum1"
	"github.com/sikalabs/tergum/tergum2"
)

func main() {
	// Backup parameters from config file
	path := flag.String("config", "", "tergum config file (json)")
	v2 := flag.Bool("v2", false, "tergum v2")

	flag.Parse()

	if *path == "" {
		log.Fatal(errors.New("tergum require config file (-config)"))
	}

	if *v2 {
		tergum2.Tergum2(*path)
	} else {
		tergum1.Tergum1(*path)
	}
}
