package source

import (
	"fmt"
	"io"

	"github.com/sikalabs/tergum/backup/source/mongo"
	"github.com/sikalabs/tergum/backup/source/mysql"
	"github.com/sikalabs/tergum/backup/source/mysql_server"
	"github.com/sikalabs/tergum/backup/source/postgres"
	"github.com/sikalabs/tergum/backup/source/single_file"
)

type Source struct {
	Mysql       *mysql.MysqlSource              `yaml:"Mysql"`
	MysqlServer *mysql_server.MysqlServerSource `yaml:"MysqlServer"`
	Postgres    *postgres.PostgresSource        `yaml:"Postgres"`
	Mongo       *mongo.MongoSource              `yaml:"Mongo"`
	SingleFile  *single_file.SingleFileSource   `yaml:"SingleFile"`
}

func (s Source) Validate() error {
	if s.Mysql != nil {
		m := *s.Mysql
		return m.Validate()
	}

	if s.MysqlServer != nil {
		m := *s.MysqlServer
		return m.Validate()
	}

	if s.Postgres != nil {
		p := *s.Postgres
		return p.Validate()
	}

	if s.Mongo != nil {
		p := *s.Mongo
		return p.Validate()
	}

	if s.SingleFile != nil {
		p := *s.SingleFile
		return p.Validate()
	}

	return fmt.Errorf("source/validate: no source detected")
}

func (s Source) Backup() (io.ReadSeeker, error) {
	if s.MysqlServer != nil {
		m := *s.MysqlServer
		return m.Backup()
	}

	if s.Mysql != nil {
		m := *s.Mysql
		return m.Backup()
	}

	if s.Postgres != nil {
		p := *s.Postgres
		return p.Backup()
	}

	if s.Mongo != nil {
		p := *s.Mongo
		return p.Backup()
	}

	if s.SingleFile != nil {
		p := *s.SingleFile
		return p.Backup()
	}

	return nil, fmt.Errorf("source/backup: no source detected")
}
