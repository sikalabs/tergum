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

__Tergum is under activedevelopment, not all features are already implemented.__

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
