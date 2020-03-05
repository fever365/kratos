package fanout

<<<<<<< HEAD
import "github.com/fever365/kratos/pkg/stat/metric"
=======
import "github.com/fever365/kratos/pkg/stat/metric"
>>>>>>> 3c6dbc7bf446fcf807931c0adeb03ddb0e59f774

const namespace = "sync"

var (
	_metricChanSize = metric.NewGaugeVec(&metric.GaugeVecOpts{
		Namespace: namespace,
		Subsystem: "pipeline_fanout",
		Name:      "current",
		Help:      "sync pipeline fanout current channel size.",
		Labels:    []string{"name"},
	})
)
