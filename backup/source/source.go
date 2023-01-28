package source

import (
	"fmt"

	"github.com/sikalabs/tergum/backup/source/dir"
	"github.com/sikalabs/tergum/backup/source/dummy"
	"github.com/sikalabs/tergum/backup/source/ftp"
	"github.com/sikalabs/tergum/backup/source/kubernetes"
	"github.com/sikalabs/tergum/backup/source/kubernetes_tls_secret"
	"github.com/sikalabs/tergum/backup/source/mongo"
	"github.com/sikalabs/tergum/backup/source/mysql"
	"github.com/sikalabs/tergum/backup/source/mysql_server"
	"github.com/sikalabs/tergum/backup/source/notion"
	"github.com/sikalabs/tergum/backup/source/postgres"
	"github.com/sikalabs/tergum/backup/source/postgres_server"
	"github.com/sikalabs/tergum/backup/source/redis"
	"github.com/sikalabs/tergum/backup/source/single_file"
	"github.com/sikalabs/tergum/backup/source/vault"
	"github.com/sikalabs/tergum/backup_output"
)

type Source struct {
	Mysql               *mysql.MysqlSource                         `yaml:"Mysql"`
	MysqlServer         *mysql_server.MysqlServerSource            `yaml:"MysqlServer"`
	Postgres            *postgres.PostgresSource                   `yaml:"Postgres"`
	PostgresServer      *postgres_server.PostgresServerSource      `yaml:"PostgresServer"`
	Mongo               *mongo.MongoSource                         `yaml:"Mongo"`
	SingleFile          *single_file.SingleFileSource              `yaml:"SingleFile"`
	KubernetesTLSSecret *kubernetes_tls_secret.KubernetesTLSSecret `yaml:"KubernetesTLSSecret"`
	Kubernetes          *kubernetes.Kubernetes                     `yaml:"Kubernetes"`
	Dir                 *dir.DirSource                             `yaml:"Dir"`
	Notion              *notion.NotionSource                       `yaml:"Notion"`
	FTP                 *ftp.FTPSource                             `yaml:"FTP"`
	Redis               *redis.RedisSource                         `yaml:"Redis"`
	Vault               *vault.VaultSource                         `yaml:"Vault"`
	Dummy               *dummy.DummySource                         `yaml:"Dummy"`
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

	if s.PostgresServer != nil {
		p := *s.PostgresServer
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

	if s.KubernetesTLSSecret != nil {
		p := *s.KubernetesTLSSecret
		return p.Validate()
	}

	if s.Kubernetes != nil {
		p := *s.Kubernetes
		return p.Validate()
	}

	if s.Dir != nil {
		p := *s.Dir
		return p.Validate()
	}

	if s.Notion != nil {
		p := *s.Notion
		return p.Validate()
	}

	if s.FTP != nil {
		p := *s.FTP
		return p.Validate()
	}

	if s.Redis != nil {
		p := *s.Redis
		return p.Validate()
	}

	if s.Vault != nil {
		p := *s.Vault
		return p.Validate()
	}

	if s.Dummy != nil {
		p := *s.Dummy
		return p.Validate()
	}

	return fmt.Errorf("source/validate: no source detected")
}

func (s Source) Backup() (backup_output.BackupOutput, error) {
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

	if s.PostgresServer != nil {
		p := *s.PostgresServer
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

	if s.KubernetesTLSSecret != nil {
		p := *s.KubernetesTLSSecret
		return p.Backup()
	}

	if s.Kubernetes != nil {
		p := *s.Kubernetes
		return p.Backup()
	}

	if s.Dir != nil {
		p := *s.Dir
		return p.Backup()
	}

	if s.Notion != nil {
		p := *s.Notion
		return p.Backup()
	}

	if s.FTP != nil {
		p := *s.FTP
		return p.Backup()
	}

	if s.Redis != nil {
		p := *s.Redis
		return p.Backup()
	}

	if s.Vault != nil {
		p := *s.Vault
		return p.Backup()
	}

	if s.Dummy != nil {
		p := *s.Dummy
		return p.Backup()
	}

	return backup_output.BackupOutput{}, fmt.Errorf("source/backup: no source detected")
}

func (s Source) Name() string {
	if s.MysqlServer != nil {
		return "MysqlServer"
	}

	if s.Mysql != nil {
		return "Mysql"
	}

	if s.Postgres != nil {
		return "Postgres"
	}

	if s.PostgresServer != nil {
		return "PostgresServer"
	}

	if s.Mongo != nil {
		return "Mongo"
	}

	if s.SingleFile != nil {
		return "SingleFile"
	}

	if s.KubernetesTLSSecret != nil {
		return "KubernetesTLSSecret"
	}

	if s.Kubernetes != nil {
		return "Kubernetes"
	}

	if s.Dir != nil {
		return "Dir"
	}

	if s.Notion != nil {
		return "Notion"
	}

	if s.FTP != nil {
		return "FTP"
	}

	if s.Redis != nil {
		return "Redis"
	}

	if s.Vault != nil {
		return "Vault"
	}

	if s.Dummy != nil {
		return "Dummy"
	}

	return ""
}
