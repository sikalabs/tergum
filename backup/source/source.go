package source

import (
	"fmt"

	"github.com/sikalabs/tergum/backup/source/consul"
	"github.com/sikalabs/tergum/backup/source/dir"
	"github.com/sikalabs/tergum/backup/source/dummy"
	"github.com/sikalabs/tergum/backup/source/ftp"
	"github.com/sikalabs/tergum/backup/source/gitlab"
	"github.com/sikalabs/tergum/backup/source/kubernetes"
	"github.com/sikalabs/tergum/backup/source/kubernetes_tls_secret"
	"github.com/sikalabs/tergum/backup/source/mongo"
	"github.com/sikalabs/tergum/backup/source/mysql"
	"github.com/sikalabs/tergum/backup/source/mysql_server"
	"github.com/sikalabs/tergum/backup/source/notion"
	"github.com/sikalabs/tergum/backup/source/postgres"
	"github.com/sikalabs/tergum/backup/source/postgres_server"
	"github.com/sikalabs/tergum/backup/source/proxmox_local_vm"
	"github.com/sikalabs/tergum/backup/source/redis"
	"github.com/sikalabs/tergum/backup/source/single_file"
	"github.com/sikalabs/tergum/backup/source/vault"
	"github.com/sikalabs/tergum/backup_output"
)

type Source struct {
	Mysql               *mysql.MysqlSource                         `yaml:"Mysql" json:"Mysql,omitempty"`
	MysqlServer         *mysql_server.MysqlServerSource            `yaml:"MysqlServer" json:"MysqlServer,omitempty"`
	Postgres            *postgres.PostgresSource                   `yaml:"Postgres" json:"Postgres,omitempty"`
	PostgresServer      *postgres_server.PostgresServerSource      `yaml:"PostgresServer" json:"PostgresServer,omitempty"`
	Mongo               *mongo.MongoSource                         `yaml:"Mongo" json:"Mongo,omitempty"`
	SingleFile          *single_file.SingleFileSource              `yaml:"SingleFile" json:"SingleFile,omitempty"`
	KubernetesTLSSecret *kubernetes_tls_secret.KubernetesTLSSecret `yaml:"KubernetesTLSSecret" json:"KubernetesTLSSecret,omitempty"`
	Kubernetes          *kubernetes.Kubernetes                     `yaml:"Kubernetes" json:"Kubernetes,omitempty"`
	Dir                 *dir.DirSource                             `yaml:"Dir" json:"Dir,omitempty"`
	Notion              *notion.NotionSource                       `yaml:"Notion" json:"Notion,omitempty"`
	FTP                 *ftp.FTPSource                             `yaml:"FTP" json:"FTP,omitempty"`
	Redis               *redis.RedisSource                         `yaml:"Redis" json:"Redis,omitempty"`
	Vault               *vault.VaultSource                         `yaml:"Vault" json:"Vault,omitempty"`
	Dummy               *dummy.DummySource                         `yaml:"Dummy" json:"Dummy,omitempty"`
	Gitlab              *gitlab.GitlabSource                       `yaml:"Gitlab" json:"Gitlab,omitempty"`
	Consul              *consul.ConsulSource                       `yaml:"Consul" json:"Consul,omitempty"`
	ProxmoxLocalVM      *proxmox_local_vm.ProxmoxLocalVMSoure      `yaml:"ProxmoxLocalVM" json:"ProxmoxLocalVM,omitempty"`
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

	if s.Gitlab != nil {
		p := *s.Gitlab
		return p.Validate()
	}

	if s.Consul != nil {
		p := *s.Consul
		return p.Validate()
	}

	if s.ProxmoxLocalVM != nil {
		p := *s.ProxmoxLocalVM
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

	if s.Gitlab != nil {
		p := *s.Gitlab
		return p.Backup()
	}

	if s.Consul != nil {
		p := *s.Consul
		return p.Backup()
	}

	if s.ProxmoxLocalVM != nil {
		p := *s.ProxmoxLocalVM
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

	if s.Gitlab != nil {
		return "Gitlab"
	}

	if s.Consul != nil {
		return "Consul"
	}

	if s.ProxmoxLocalVM != nil {
		return "ProxmoxLocalVM"
	}

	return ""
}
