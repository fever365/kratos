package main

import (
	"os/exec"

	"github.com/urfave/cli"
)

const (
<<<<<<< HEAD
	_getSwaggerGen = "go get -u github.com/bilibili/kratos/tool/protobuf/protoc-gen-bswagger"
=======
	_getSwaggerGen = "go get -u github.com/fever365/kratos/tool/protobuf/protoc-gen-bswagger"
>>>>>>> 3c6dbc7bf446fcf807931c0adeb03ddb0e59f774
	_swaggerProtoc = "protoc --proto_path=%s --proto_path=%s --proto_path=%s --bswagger_out=:."
)

func installSwaggerGen() error {
	if _, err := exec.LookPath("protoc-gen-bswagger"); err != nil {
		if err := goget(_getSwaggerGen); err != nil {
			return err
		}
	}
	return nil
}

func genSwagger(ctx *cli.Context) error {
	return generate(ctx, _swaggerProtoc)
}
