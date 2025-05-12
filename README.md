<div align="center">

<img src="readme/pull-push-transparent.png" width="400px">

# PullPush

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/mjarmoc/nx-s3-server?style=flat)
![GitHub forks](https://img.shields.io/github/forks/mjarmoc/nx-s3-server?style=flat)

Pull Push is a small application helping you to move files across different Storages.
Under the hood it utilize parallel multipart upload/download to move things fast.

</div>

## Features

- [x] Http => S3
- [ ] Http => GPC Storage
- [ ] Http => Azure Storage
- [ ] Azure Storage <=> GPC Storage
- [ ] S3 <=> Azure Storage
- [ ] S3 <=> GPC Storage

## Installation

```sh
go mod download
```

## Local Development

With localstack:

```sh
docker-compose up
```
