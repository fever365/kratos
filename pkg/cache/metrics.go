package cache

<<<<<<< HEAD
import "github.com/bilibili/kratos/pkg/stat/metric"
=======
import "github.com/fever365/kratos/pkg/stat/metric"
>>>>>>> 3c6dbc7bf446fcf807931c0adeb03ddb0e59f774

const _metricNamespace = "cache"

// be used in tool/kratos-gen-bts
var (
	MetricHits = metric.NewCounterVec(&metric.CounterVecOpts{
		Namespace: _metricNamespace,
		Subsystem: "",
		Name:      "hits_total",
		Help:      "cache hits total.",
		Labels:    []string{"name"},
	})
	MetricMisses = metric.NewCounterVec(&metric.CounterVecOpts{
		Namespace: _metricNamespace,
		Subsystem: "",
		Name:      "misses_total",
		Help:      "cache misses total.",
		Labels:    []string{"name"},
	})
)
