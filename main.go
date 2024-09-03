package main

import (
	"fmt"
	"math/rand/v2"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/razzaghis/prometheus-up-and-running/internal/metric"
)

func handler(w http.ResponseWriter, r *http.Request) {
	metric.HelloWorldCounter.Inc()
	euros := rand.IntN(10)
	metric.SalesCounter.Add(float64(euros))
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Hello World for %d euros", euros)))
}

func main() {

	// Start HTTP server
	http.Handle("/metrics", promhttp.Handler())

	http.HandleFunc("/", handler)
	fmt.Println("Starting server at :8000")
	http.ListenAndServe(":8000", nil)
}
