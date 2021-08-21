package backup

import (
	"fmt"

	"github.com/sikalabs/tergum/backup/middleware"
	"github.com/sikalabs/tergum/backup/source"
	"github.com/sikalabs/tergum/backup/target"
)

type Backup struct {
	ID          string
	Source      *source.Source
	Middlewares []middleware.Middleware
	Targets     []target.Target
}

func (b Backup) Validate() error {
	// Validate Source
	err := b.Source.Validate()
	if err != nil {
		return err
	}

	// Validate all Middlewares
	for _, m := range b.Middlewares {
		err = m.Validate()
		if err != nil {
			return err
		}
	}

	// Must have at least one Target
	if len(b.Targets) == 0 {
		return fmt.Errorf("no targets defined")
	}

	// Validate all Targets
	for _, t := range b.Targets {
		err = t.Validate()
		if err != nil {
			return err
		}
	}

	return nil
}
