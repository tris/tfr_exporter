package main

import (
	"log"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	defaultPort = "9866" // 9 + "FAA" in beeper code, though this conflicts with rtrmon exporter
)

func main() {
	http.HandleFunc("/scrape", scrapeHandler)
	http.Handle("/metrics", promhttp.Handler()) // default Go metrics

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
