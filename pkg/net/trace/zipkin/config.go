package zipkin

import (
	"time"

<<<<<<< HEAD
	"github.com/fever365/kratos/pkg/conf/env"
	"github.com/fever365/kratos/pkg/net/trace"
	xtime "github.com/fever365/kratos/pkg/time"
=======
	"github.com/fever365/kratos/pkg/conf/env"
	"github.com/fever365/kratos/pkg/net/trace"
	xtime "github.com/fever365/kratos/pkg/time"
>>>>>>> 3c6dbc7bf446fcf807931c0adeb03ddb0e59f774
)

// Config config.
// url should be the endpoint to send the spans to, e.g.
// http://localhost:9411/api/v2/spans
type Config struct {
	Endpoint      string         `dsn:"endpoint"`
	BatchSize     int            `dsn:"query.batch_size,100"`
	Timeout       xtime.Duration `dsn:"query.timeout,200ms"`
	DisableSample bool           `dsn:"query.disable_sample"`
}

// Init init trace report.
func Init(c *Config) {
	if c.BatchSize == 0 {
		c.BatchSize = 100
	}
	if c.Timeout == 0 {
		c.Timeout = xtime.Duration(200 * time.Millisecond)
	}
	trace.SetGlobalTracer(trace.NewTracer(env.AppID, newReport(c), c.DisableSample))
}
