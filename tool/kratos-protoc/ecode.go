package main

import (
	"os/exec"

	"github.com/urfave/cli"
)

const (
<<<<<<< HEAD
	_getEcodeGen = "go get -u github.com/fever365/kratos/tool/protobuf/protoc-gen-ecode"
=======
	_getEcodeGen = "go get -u github.com/fever365/kratos/tool/protobuf/protoc-gen-ecode"
>>>>>>> 3c6dbc7bf446fcf807931c0adeb03ddb0e59f774
	_ecodeProtoc = "protoc --proto_path=%s --proto_path=%s --proto_path=%s --ecode_out=:."
)

func installEcodeGen() error {
	if _, err := exec.LookPath("protoc-gen-ecode"); err != nil {
		if err := goget(_getEcodeGen); err != nil {
			return err
		}
	}
	return nil
}

func genEcode(ctx *cli.Context) error {
	return generate(ctx, _ecodeProtoc)
}
