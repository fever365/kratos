module github.com/bilibili/kratos

go 1.12

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/StackExchange/wmi v0.0.0-20190523213609-cbe66965904d // indirect
	github.com/bgentry/speakeasy v0.1.0 // indirect
	github.com/cockroachdb/datadriven v0.0.0-20190531201743-edce55837238 // indirect
	github.com/dgryski/go-farm v0.0.0-20190423205320-6a90982ecee2
	github.com/dustin/go-humanize v0.0.0-20171111073723-bb3d318650d4 // indirect
	github.com/fatih/color v1.7.0
	github.com/fever365/kratos v0.2.13
	github.com/fsnotify/fsnotify v1.4.7
	github.com/go-sql-driver/mysql v1.4.1
	github.com/gogo/protobuf v1.2.2-0.20190730201129-28a6bbf47e48
	github.com/golang/protobuf v1.3.2
	github.com/grpc-ecosystem/go-grpc-middleware v1.0.1-0.20190118093823-f849b5445de4 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/konsorten/go-windows-terminal-sequences v1.0.2 // indirect
	github.com/mattn/go-runewidth v0.0.2 // indirect
	github.com/montanaflynn/stats v0.5.0
	github.com/olekukonko/tablewriter v0.0.0-20170122224234-a0225b3f23b5 // indirect
	github.com/openzipkin/zipkin-go v0.2.1
	github.com/pkg/errors v0.8.1
	github.com/prometheus/client_golang v1.1.0
	github.com/prometheus/client_model v0.0.0-20190220174349-fd36f4220a90 // indirect
	github.com/remyoudompheng/bigfft v0.0.0-20190806203942-babf20351dd7e3ac320adedbbe5eb311aec8763c // indirect
	github.com/shirou/gopsutil v2.19.6+incompatible
	github.com/siddontang/go v0.0.0-20180604090527-bdc77568d726
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/cobra v0.0.3 // indirect
	github.com/spf13/pflag v1.0.1 // indirect
	github.com/stretchr/testify v1.4.0
	github.com/tsuna/gohbase v0.0.0-20190823190353-a66bcc9075db
	github.com/urfave/cli v1.21.0
	go.etcd.io/etcd v3.3.15+incompatible
	golang.org/x/net v0.0.0-20190827160401-ba9fcec4b297
	golang.org/x/time v0.0.0-20190513212739-9d24e82272b4 // indirect
	google.golang.org/genproto v0.0.0-20190819201941-24fa4b261c55
	google.golang.org/grpc v1.23.0
	gopkg.in/airbrake/gobrake.v2 v2.0.9 // indirect
	gopkg.in/cheggaaa/pb.v1 v1.0.25 // indirect
	gopkg.in/gemnasium/logrus-airbrake-hook.v2 v2.1.2 // indirect
	gopkg.in/go-playground/validator.v9 v9.29.1
	gopkg.in/yaml.v2 v2.2.2
	sigs.k8s.io/yaml v1.1.0 // indirect
)

replace github.com/bilibili/kratos v0.2.3 => github.com/fever365/kratos v0.2.13
