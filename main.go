package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/razzaghis/prometheus-up-and-running/internal/metric"
)

func handler(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	metric.InProgressGauge.Inc()
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello World"))
	metric.LastServedGauge.SetToCurrentTime()
	metric.InProgressGauge.Dec()
	metric.LatencyHelloWorld.Observe(time.Since(startTime).Seconds())
}

func main() {

	// Start HTTP server
	http.Handle("/metrics", promhttp.Handler())

	http.HandleFunc("/", handler)
	fmt.Println("Starting server at :8000")
	http.ListenAndServe(":8000", nil)
}
