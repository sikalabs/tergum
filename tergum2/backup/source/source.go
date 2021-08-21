package source

import (
	"fmt"

	"github.com/sikalabs/tergum/tergum2/backup/source/mysql"
	"github.com/sikalabs/tergum/tergum2/backup/source/postgres"
)

type Source struct {
	Mysql    *mysql.MysqlSource
	Postgres *postgres.PostgresSource
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
