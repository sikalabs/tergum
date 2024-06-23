package backup

import (
	"fmt"

	"github.com/sikalabs/tergum/backup/middleware"
	"github.com/sikalabs/tergum/backup/source"
	"github.com/sikalabs/tergum/backup/target"
)

type RemoteExec struct {
	Server string `yaml:"Server" json:"Server,omitempty"`
}

type Backup struct {
	ID          string                  `yaml:"ID" json:"ID,omitempty"`
	RemoteExec  *RemoteExec             `yaml:"RemoteExec" json:"RemoteExec,omitempty"`
	Source      *source.Source          `yaml:"Source" json:"Source,omitempty"`
	Middlewares []middleware.Middleware `yaml:"Middlewares" json:"Middlewares,omitempty"`
	Targets     []target.Target         `yaml:"Targets" json:"Targets,omitempty"`
	SleepBefore int                     `yaml:"SleepBefore" json:"SleepBefore,omitempty"`
}

func (b Backup) Validate() error {
	// Validate Source
	if b.Source == nil {
		return fmt.Errorf("backup/validate: source is not defined")
	}
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
