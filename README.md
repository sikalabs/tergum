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

## What "Tergum" means?

Tergum means backup in latin.

## Install

Install using Brew:

```
brew install sikalabs/tap/tergum
```

## Usage

Tergum has only one CLI argumet which points to config file.

```
tergum -config tergum.json
```

### Tergum Config File

Tergum supports only JSON config file, but we're working for YAML support.

Config file examples are in [misc/example/config](./misc/example/config) directory

#### Basic Config Structure

```jsx
{
  "meta": {
    "schemaVersion": 2
  },
  "alerting": <Alerting>,
  "backups": [
    <Backup>,
    <Backup>,
    ...
  ]
}
```

#### Backup Block

```jsx
{
  "id": <UniqueBackupID>,
  "source": {
    "name": <BackupSourceBackend (mysq,)>,
    "mysql": <BackupSourceMysqlConfiguration>,
  },
  "destinations": [
    {
      "id": <UniqueBackupDestinationID>,
      "name": <BackupDestinationBackend (filepath, file, s3)>,
      "filePath": <BackupDestinationFilePathConfiguration>,
      "file": <BackupDestinationFileConfiguration>,
      "s3": <BackupDestinationS3Configuration>,
    },
    ...
  ]
}
```

#### Example BackupSourceMysqlConfiguration Block

```jsx
{
  "host": "127.0.0.1",
  "port": "3306",
  "user": "root",
  "password": "root",
  "database": "default"
}
```


#### Example BackupDestinationFilePathConfiguration Block

```jsx
{
  "path": "/backup/mysql-default.sql"
}
```

#### Example BackupDestinationFileConfiguration Block

```jsx
{
  "dir": "/backup/",
  "prefix": "mysql-default",
  "suffix": "sql"
}
```

#### Example BackupDestinationS3Configuration Block

AWS:

```jsx
{
  "accessKey": "admin",
  "secretKey": "asdfasdf",
  "endpoint": "https://minio.example.com",
  "bucketName": "tergum-backups",
  "prefix": "mysql-default",
  "suffix": "sql"
}
```

Minio:

```jsx
{
  "accessKey": "aws_access_key_id",
  "secretKey": "aws_secret_access_key",
  "region": "eu-central-1",
  "bucketName": "tergum-backups",
  "prefix": "mysql-default",
  "suffix": "sql"
}
```

#### Alerting Block

```jsx
{
  "Backends": {
    "Email":  <AlertingBackendEmail>
  },
  "Alerts":[
    <Alert>,
    <Alert>,
    ...
  ]
}
```

#### Example AlertingBackendEmail Block

```jsx
{
  "smtpHost": "mail.example.com",
  "smtpPort": "25",
  "email": "tergum@example.com",
  "password": "asdfasdf"
}
```

#### Alert Block

```jsx
{
  "Backend": <AlertBackendName (email,)>,
  "Email": <AlertEmail>,
}
```


#### Example AlertEmail Block

```jsx
{
  "Emails": [
    "ondrej@example.com",
    "monitoring@example.com"
  ]
}
```

## Current Project State

### Backup Sources

- [ ] Files
- [ ] Postgres
- [x] MySQL
- [ ] S3

### Backup Processors

- [ ] ZIP Compression
- [ ] Symetric Encryption
- [ ] Asymetric Encryption
- [ ] GPG Encryption
- [ ] GPG Signatures

### Backup Storage

- [x] Files
- [x] S3

### Alerting

- [x] Email
- [ ] Slack
- [ ] Pagerduty
