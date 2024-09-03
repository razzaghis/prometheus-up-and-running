// Package metric is the package for prometheus metrics
package metric

import (
	"github.com/prometheus/client_golang/prometheus"
)

func init() {
	prometheus.MustRegister(HelloWorldCounter)
}

// HelloWorldCounter is used for recording the count of hello world messages
var HelloWorldCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Namespace: "ns",
		Subsystem: "test",
		Name:      "hello_world_total",
		Help:      "Hello Worlds requested.",
	},
)
