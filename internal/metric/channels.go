// Package metric is the package for prometheus metrics
package metric

import (
	"github.com/prometheus/client_golang/prometheus"
)

func init() {
	prometheus.MustRegister(HelloWorldCounter)
	prometheus.MustRegister(SalesCounter)
	prometheus.MustRegister(InProgressGauge)
	prometheus.MustRegister(LastServedGauge)
	prometheus.MustRegister(LatencyHelloWorld)
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

// InProgressGauge is used for recording the in progress hello world requests
var InProgressGauge = prometheus.NewGauge(
	prometheus.GaugeOpts{
		Namespace: "ns",
		Subsystem: "test",
		Name:      "hello_worlds_in_progress",
		Help:      "Number of Hello Worlds in progress.",
	},
)

// LastServedGauge is used for recording the last time a Hello World was served
var LastServedGauge = prometheus.NewGauge(
	prometheus.GaugeOpts{
		Namespace: "ns",
		Subsystem: "test",
		Name:      "hello_world_last_time_seconds",
		Help:      "The last time a Hello World was served.",
	},
)

// LatencyHelloWorld is used for recording the latency if Hello World requests
var LatencyHelloWorld = prometheus.NewHistogram(
	prometheus.HistogramOpts{
		Namespace: "ns",
		Subsystem: "test",
		Name:      "hello_world_latency_seconds",
		Help:      "Time for a request Hello World.",
	},
)
