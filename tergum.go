package main

import (
	"errors"
	"flag"
	"fmt"
	"log"

	"github.com/sikalabs/tergum/tergum1"
	"github.com/sikalabs/tergum/tergum2"
	"github.com/sikalabs/tergum/version"
)

func main() {
	// Backup parameters from config file
	path := flag.String("config", "", "tergum config file (json)")
	v2 := flag.Bool("v2", false, "tergum v2")
	ver := flag.Bool("v", false, "show tergum version")

	flag.Parse()

	if *ver {
		fmt.Println(version.Version)
		return
	}

	if *path == "" {
		log.Fatal(errors.New("tergum require config file (-config)"))
	}

	if *v2 {
		tergum2.Tergum2(*path)
	} else {
		tergum1.Tergum1(*path)
	}
}
