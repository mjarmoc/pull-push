package main

import (
	"github.com/mjarmoc/pull-push/cmd"
	"github.com/mjarmoc/pull-push/cmd/gcp2s3"
)

func main() {
	gcp2s3.Init()
	cmd.Execute()
}
