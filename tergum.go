package main

import (
	"errors"
	"flag"
	"fmt"
	"log"

	"github.com/sikalabs/tergum/do_backup"
	"github.com/sikalabs/tergum/src1"
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
		do_backup.DoBackup(*path)
	} else {
		src1.Tergum1(*path)
	}
}
