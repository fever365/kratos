package main

import (
	"flag"
	"fmt"
	"os"

<<<<<<< HEAD
	"github.com/fever365/kratos/tool/protobuf/pkg/gen"
	"github.com/fever365/kratos/tool/protobuf/pkg/generator"
	ecodegen "github.com/fever365/kratos/tool/protobuf/protoc-gen-ecode/generator"
=======
	"github.com/fever365/kratos/tool/protobuf/pkg/gen"
	"github.com/fever365/kratos/tool/protobuf/pkg/generator"
	ecodegen "github.com/fever365/kratos/tool/protobuf/protoc-gen-ecode/generator"
>>>>>>> 3c6dbc7bf446fcf807931c0adeb03ddb0e59f774
)

func main() {
	versionFlag := flag.Bool("version", false, "print version and exit")
	flag.Parse()
	if *versionFlag {
		fmt.Println(generator.Version)
		os.Exit(0)
	}

	g := ecodegen.EcodeGenerator()
	gen.Main(g)
}
