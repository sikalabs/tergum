<p align="center">
  <h1 align="center">Tergum: Universal Backup Tool</h1>
  <p align="center">
    <a href="https://opensource.sikalabs.com"><img alt="SikaLabs" src="https://img.shields.io/badge/OPENSOURCE BY-SIKALABS-131480?style=for-the-badge"></a>
    <a href="https://sikalabs.com"><img alt="SikaLabs" src="https://img.shields.io/badge/-sikalabs.com-gray?style=for-the-badge"></a>
    <a href="mailto://opensource@sikalabs.com"><img alt="SikaLabs" src="https://img.shields.io/badge/-opensource@sikalabs.com-gray?style=for-the-badge"></a>
  </p>
</p>

## Why Tergum?

Tergum is simple tool provides centralized backup solution with multiple sources (databases, files, S3, ...) and multiple backup storages (S3, filesystem, ...). Tergum has native backup monitoring and alerts you when backup fails. Tergum also support backup encryption, compression and automatic recovery testing.

__Tergum is under active development, not all features are already implemented. Check [current project state](#current-project-state)__

## Do you want to start using Tergum? Give us a call

Let's discuss Tergum in your project in [30 min call](https://calendly.com/ondrejsika/tergum-intro)

## What "Tergum" means?

Tergum means backup in latin.

## Tergum Cloud: Bring Your Backups into Cloud

Tergum Cloud allow you to manage your backup using UI & Terraform and store your backups securely in our AWS.

Are you interested in our public beta? Drop us email [hello@sikalabs.com](mailto:hello@sikalabs.com?subject=Tergum_Cloud)

## Tergum Enterprise: Use Tergum Cloud in Your Private Infrastructure

Tergum Enterprise brings our cloud platform behind your filewall. For an inquiry, contact our sales [sales@sikalabs.com](mailto:sales@sikalabs.com?subject=Tergum_Enterprise)

## Install

Install using Brew:

```
brew install sikalabs/tap/tergum
```

### Autocomplete

See: `tergum completion`

#### Bash

```
source <(tergum completion bash)
```

## CLI Usage

### Generated CLI Docs on Github

See: <https://github.com/sikalabs/tergum-cli-docs/blob/master/tergum.md#tergum>

## Generate CLI Docs

Generate Markdown CLI docs to `./cobra-docs`

```
tergum generate-docs
```

## Tergum Config File

Tergum supports only JSON config file, but we're working on YAML support.

Config file examples are in [misc/example/config](./misc/example/config) directory

#### Basic Config Structure

```yaml
Meta:
  SchemaVersion: 3
Cloud: <Cloud>
Notification: <Notification>
Backups:
  - <Backup>
  - <Backup>
  - ...
```

#### Backup Block

```yaml
ID: <UniqueBackupID>
Source:
  Mysql: <BackupSourceMysqlConfiguration>
  MysqlServer: <BackupSourceMysqlServerConfiguration>
  Postgres: <BackupSourcePostgresConfiguration>
  PostgresServer: <BackupSourcePostgresServerConfiguration>
  Mongo: <BackupSourceMongoConfiguration>
  SingleFile: <BackupSourceSingleFileConfiguration>
  Dir: <BackupSourceDirConfiguration>
  KubernetesTLSSecret: <BackupSourceKubernetesTLSSecret>
  Kubernetes: <BackupSourceKubernetes>
  Notion: <BackupSourceNotion>
  FTP: <BackupSourceFTP>
  Redis: <BackupSourceRedis>
  Vault: <BackupSourceVault>
  Dummy: <BackupSourceDummy>
  Gitlab: <BackupSourceGitlab>
  Consul: <BackupSourceConsul>
Middlewares:
  - <MiddlewareConfiguration>
  - ...
Destinations:
  - ID: <UniqueBackupDestinationID>
    Middlewares:
      - <MiddlewareConfiguration>
      - ...
    FilePath: <BackupDestinationFilePathConfiguration>
    File: <BackupDestinationFileConfiguration>
    S3: <BackupDestinationS3Configuration>
    AzureBlob: <BackupDestinationAzureBlobConfiguration>
  - ...
SleepBefore: <sleep time befor backup job in seconds>
```

#### GzipMiddlewareConfiguration

```yaml
Gzip: {}
SymmetricEncryption:
  Passphrase: "passphrase"
```

#### Example BackupSourceMysqlConfiguration Block

```yaml
Host: "127.0.0.1"
Port: "3306"
User: "root"
Password: "root"
Database: "default"
```

With extra args

```yaml
Host: "127.0.0.1"
Port: "3306"
User: "root"
Password: "root"
Database: "default"
MysqldumpExtraArgs:
  - --column-statistics=0
```

#### Example BackupSourceMysqlServerConfiguration Block

```yaml
Host: "127.0.0.1"
Port: "3306"
User: "root"
Password: "root"
```

With extra args

```yaml
Host: "127.0.0.1"
Port: "3306"
User: "root"
Password: "root"
MysqldumpExtraArgs:
  - --column-statistics=0
```

#### Example BackupSourcePostgresConfiguration Block

```yaml
Host: "127.0.0.1"
Port: "15432"
User: "postgres"
Password: "pg"
Database: "postgres"
```

With extra args

```yaml
Host: "127.0.0.1"
Port: "15432"
User: "postgres"
Password: "pg"
Database: "postgres"
PgdumpExtraArgs:
  - --ignore-version
```

With SSL mode

```yaml
Host: "127.0.0.1"
Port: "15432"
User: "postgres"
Password: "pg"
Database: "postgres"
SSLMode: "require"
```

#### Example BackupSourcePostgresServerConfiguration Block

```yaml
Host: "127.0.0.1"
Port: "15432"
User: "postgres"
Password: "pg"
```

With extra args

```yaml
Host: "127.0.0.1"
Port: "15432"
User: "postgres"
Password: "pg"
PgdumpallExtraArgs:
  - --ignore-version
```

With SSL mode

```yaml
Host: "127.0.0.1"
Port: "15432"
User: "postgres"
Password: "pg"
SSLMode: "require"
```

#### Example BackupSourceMongoConfiguration Block

Dump all dbs & no auth

```yaml
Host: "127.0.0.1"
Port: "27017"
```

Dump all dbs with auth

```yaml
Host: "127.0.0.1"
Port: "27017"
User: "root"
Password: "root"
```

Dump single db with auth

```yaml
Host: "127.0.0.1"
Port: "27017"
User: "root"
Password: "root"
Database: "test"
```

Dump single db with auth and custom Authentication Database

```yaml
Host: "127.0.0.1"
Port: "27017"
User: "root"
Password: "root"
AuthenticationDatabase: "test" # default is admin
Database: "test"
```

#### Example BackupSourceKubernetesTLSSecret Block

Backup all TLS secrets

```yaml
Server: https://kubernetes-api.example.com
Token: foo-bar-baz
Namespace: default
```

Backup single TLS secret

```yaml
Server: https://kubernetes-api.example.com
Token: foo-bar-baz
Namespace: default
SecretName: tls-example-com
```

#### Example BackupSourceKubernetes Block

Backup all resources (pods)

```yaml
Server: https://kubernetes-api.example.com
Token: foo-bar-baz
Namespace: default
Resource: pod
```

Backup single resource (hello-world pod)

```yaml
Server: https://kubernetes-api.example.com
Token: foo-bar-baz
Namespace: default
Resource: pod
Name: hello-world
```

#### Example BackupSourceSingleFileConfiguration Block

```yaml
Path: /data/export/dump.sql
```

### Example BackupSourceDirConfiguration Block

```yaml
Path: /data
Excludes:
  - /data/tmp
```


### Example BackupSourceNotion Block

```yaml
Token: <Notion token_v2>
SpaceID: <Notion Space UID>
Format: <Fotmat of export ("html" or "markdown")>
```

### Example BackupSourceFTP Block

```yaml
Host: <FTP host>
User: <FTP user>
Password: <FTP password>
```

### Example BackupSourceRedis Block

```yaml
Host: <host>
Port: <port>
```

### Example BackupSourceVault Block

```yaml
Addr: <vault address>
Token: <vault token>
Headers: <map[string]string of headers, optional>
```

example with cloudflare access headers

```yaml
Addr: https://vault.corp.com
Token: s.1234567890
Headers:
  CF-Access-Client-ID: xxx1234567890
  CF-Access-Client-Secret: xxx123456789
```


### Example BackupSourceDummy Block

```yaml
Content: <backup content>
```

### Example BackupSourceGitlab Block

```yaml
NamePrefix: <prefix Gitlab backup file in /var/opt/gitlab/backups>
Skip: <skip (for example registry)>
```

- Gitlab Docs about SKIP - <https://docs.gitlab.com/ee/administration/backup_restore/backup_gitlab.html?tab=Linux+package+%28Omnibus%29#excluding-specific-data-from-the-backup>

### Example BackupSourceConsul Block

```yaml
Addr: <host>
Token: <token>
```

Example without ACL

```yaml
Addr: http://127.0.0.1:8500
```

Example with ACL requires token

```yaml
Addr: http://127.0.0.1:8500
Token: 51047cd1-c243-a969-2bf1-a845405e4da9
```

#### Example BackupDestinationFilePathConfiguration Block

```yaml
Path: "/backup/mysql-default.sql"
```

#### Example BackupDestinationFileConfiguration Block

```yaml
Dir: "/backup/"
Prefix: "mysql-default"
Suffix: "sql"
```

#### Example BackupDestinationS3Configuration Block

AWS:

```yaml
AccessKey: "admin"
SecretKey: "asdfasdf"
Endpoint: "https://minio.example.com"
BucketName: "tergum-backups"
Prefix: "mysql-default"
Suffix: "sql"
```

Minio:

```yaml
accessKey: "aws_access_key_id"
secretKey: "aws_secret_access_key"
region: "eu-central-1"
bucketName: "tergum-backups"
prefix: "mysql-default"
suffix: "sql"
```

#### Example BackupDestinationAzureBlobConfiguration Block

```yaml
AccountName: account_name
AccountKey: account_key
ContainerName: container_name
Prefix: "mysql-default"
Suffix: "sql"
```

#### Notification Block

```yaml
Backends: {
  Email:  <NotificationBackendEmail>
Target:
  - <NotificationTarget>
  - <NotificationTarget>
  - ...
```

#### Example NotificationBackendEmail Block

```yaml
SmtpHost: "mail.example.com"
SmtpPort: "25"
Usename: "aaa"
Password: "aaa/bbb"
From: "tergum@example.com"
```

#### NotificationTarget Block

```yaml
Email: <NotificationEmailTarget>
SlackWebhook: <NotificationSlackWebhookTarget>
```

#### Example NotificationEmailTarget Block

```yaml
Emails:
  - ondrej@example.com
  - monitoring@example.com
SendOK: false
```

- `SendOK=true` will send email notification for all tergum runs (failed & OK runs)

#### Example NotificationSlackWebhookTarget Block

```yaml
URLs:
  - https://hooks.slack.com/services/xxx/yyy/zzz
SendOK: false
```

- `SendOK=true` will send email notification for all tergum runs (failed & OK runs)

#### Cloud Block

```yaml
Email: <email of tergum cloud account>
```

### Tergum Utils

#### `tergum utils cron`

Simple cron scheduler in Tergum

```
tergum utils cron <cron-expression> <command> [args...]
```

Example usage:

```
tergum utils cron "0 0 * * *" -- tergum backup -c tergum.yml
```

## Current Project State

### Backup Sources

- [x] SingleFile
- [x] Files (Dir)
- [x] Postgres
- [x] PostgresServer
- [x] MySQL
- [x] MySQLServer
- [ ] Oracle (Enterprise)
- [ ] S3
- [ ] Ceph RBD
- [ ] CephFS
- [x] MongoDB
- [x] Gitlab
- [ ] Proxmox
- [x] Kubernetes Resource
  - [x] Kubernetes TLS Secret
- [ ] Container Image
- [x] Redis
- [x] [Notion](https://notion.so)
- [x] FTP Server (for old school hostings)
- [x] Hashicorp Vault
- [x] Hashicorp Consul
- [x] Dummy (for testing)

### Passwords Sources

- [x] YAML
- [x] Environment Variables
- [ ] Hashicorp Vault
- [ ] AWS Secrets Manager
- [ ] Azure Key Vault

### Backup Processors

- [x] GZIP Compression
- [x] Symmetric Encryption
- [ ] AsymmetricEncryption
- [ ] GPG Encryption
- [ ] GPG Signatures

### Backup Storage

- [x] Files
- [x] S3
- [ ] Tergum Cloud
- [x] Azure Blob
- [ ] GCS (Google Cloud Storage)
- [ ] Container Registry

### Notification

- [x] Email
- [x] Slack
- [ ] Microsoft Teams
- [ ] Pagerduty
