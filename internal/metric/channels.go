// Package metric is the package for prometheus metrics
package metric

import (
	"github.com/prometheus/client_golang/prometheus"
)

func init() {
	prometheus.MustRegister(HelloWorldCounter)
	prometheus.MustRegister(SalesCounter)
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

// SalesCounter is used for recording the count of sales
var SalesCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Namespace: "ns",
		Subsystem: "test",
		Name:      "hello_world_sales_euro_total",
		Help:      "Euros made serving Hello World.",
	},
)
