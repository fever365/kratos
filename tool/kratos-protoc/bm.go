package main

import (
	"os/exec"

	"github.com/urfave/cli"
)

const (
<<<<<<< HEAD
	_getBMGen = "go get -u github.com/bilibili/kratos/tool/protobuf/protoc-gen-bm"
=======
	_getBMGen = "go get -u github.com/fever365/kratos/tool/protobuf/protoc-gen-bm"
>>>>>>> 3c6dbc7bf446fcf807931c0adeb03ddb0e59f774
	_bmProtoc = "protoc --proto_path=%s --proto_path=%s --proto_path=%s --bm_out=:."
)

func installBMGen() error {
	if _, err := exec.LookPath("protoc-gen-bm"); err != nil {
		if err := goget(_getBMGen); err != nil {
			return err
		}
	}
	return nil
}

func genBM(ctx *cli.Context) error {
	return generate(ctx, _bmProtoc)
}
