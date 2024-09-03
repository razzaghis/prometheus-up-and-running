package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/razzaghis/prometheus-up-and-running/internal/metric"
)

func handler(w http.ResponseWriter, r *http.Request) {
	metric.InProgressGauge.Inc()
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello World"))
	metric.LastServedGauge.SetToCurrentTime()
	// add sleep time to see the in progress metrics
	time.Sleep(10 * time.Second)
	metric.InProgressGauge.Dec()
}

func main() {

	// Start HTTP server
	http.Handle("/metrics", promhttp.Handler())

	http.HandleFunc("/", handler)
	fmt.Println("Starting server at :8000")
	http.ListenAndServe(":8000", nil)
}
