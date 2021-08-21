package tergum2

import (
	"fmt"

	"github.com/sikalabs/tergum/tergum2/config"
)

func Tergum2(configPath string) {
	fmt.Println("tergum v2")

	// Load config from file
	var config config.TergumConfig
	config.Load(configPath)

	// Validate config
	err := config.Validate()
	if err != nil {
		fmt.Println(err)
	}

	for _, b := range config.Backups {
		// Backup source
		data, err := b.Source.Backup()
		if err != nil {
			fmt.Println(err)
		}

		// Process Backup's Middlewares
		for _, m := range b.Middlewares {
			data, err = m.Process(data)
			if err != nil {
				fmt.Println(err)
			}
		}

		for _, t := range b.Targets {
			targetData := data

			// Process Targets's Middlewares
			for _, m := range t.Middlewares {
				targetData, err = m.Process(targetData)
				if err != nil {
					fmt.Println(err)
				}
			}

			// Save backup to target
			err = t.Save(targetData)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
