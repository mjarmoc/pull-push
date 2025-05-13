<div align="center">

<img src="readme/pull-push-transparent.png" width="400px">

# Pull-Push

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/mjarmoc/pull-push?style=flat)
![GitHub releases](https://img.shields.io/github/release-date/mjarmoc/pull-push?style=flat)
![GitHub forks](https://img.shields.io/github/forks/mjarmoc/pull-push?style=flat)
![GitHub downloads](https://img.shields.io/github/downloads/mjarmoc/pull-push/total?style=flat)

Pull-Push is a small application helping you to move files across different Storages.<br/>
Under the hood it utilize parallel multipart upload/download to move things fast.

</div>

## Features

- [x] Http => S3
- [ ] Http => GPC Storage
- [ ] Http => Azure Storage
- [ ] Azure Storage => GCP Storage
- [ ] Azure Storage => S3
- [ ] S3 => Azure Storage
- [ ] S3 => GCP Storage
- [ ] GCP Storage => S3 [in progress]
- [ ] GCP Storage => Azure Storage

## Installation

```sh
go install
```

## Usage

Command flags are not stable at the moment and will change.

Http -> S3

```sh
pull-push -u <file-url> -b <bucket-name> -f <file-path>
```

## Local Development

```sh
go mod download
```

With localstack:

```sh
docker-compose up
```
