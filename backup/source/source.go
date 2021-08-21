package source

import (
	"fmt"

	"github.com/sikalabs/tergum/backup/source/mysql"
	"github.com/sikalabs/tergum/backup/source/postgres"
)

type Source struct {
	Mysql    *mysql.MysqlSource       `yaml:"Mysql"`
	Postgres *postgres.PostgresSource `yaml:"Postgres"`
}

func (s Source) Validate() error {
	if s.Mysql != nil {
		m := *s.Mysql
		return m.Validate()
	}

	if s.Postgres != nil {
		p := *s.Postgres
		return p.Validate()
	}

	return fmt.Errorf("no source detected")
}

func (s Source) Backup() ([]byte, error) {
	if s.Mysql != nil {
		m := *s.Mysql
		return m.Backup()
	}

	if s.Postgres != nil {
		p := *s.Postgres
		return p.Backup()
	}

	return nil, fmt.Errorf("no source detected")
}
