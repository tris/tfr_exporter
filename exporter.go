package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	token = ""

	labelNames = []string{
		"latitude",
		"longitude",
	}
)

func init() {
	token = os.Getenv("ALOFT_TOKEN")
	if token == "" {
		log.Fatal("ALOFT_TOKEN environment variable is required")
	}
}

func scrapeHandler(w http.ResponseWriter, r *http.Request) {
	latitude := r.URL.Query().Get("lat")
	if latitude == "" {
		http.Error(w, "Missing required parameter: lat", http.StatusBadRequest)
		return
	}
	longitude := r.URL.Query().Get("long")
	if longitude == "" {
		http.Error(w, "Missing required parameter: long", http.StatusBadRequest)
		return
	}

	// Fetch current advisories from Aloft
	// API docs: https://api.aloft.ai/v1/docs
	req, err := http.NewRequest("GET", "https://api.aloft.ai/v1/airspace/advisories?advisory_types%5B%5D=faa_tfrs&lat="+latitude+"&lng="+longitude, nil)
	req.Header.Add("Authorization", "Bearer "+token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var airspace AirspaceResponse
	if err := json.Unmarshal(body, &airspace); err != nil {
		log.Println(err)
		return
	}

	if !airspace.Success {
		log.Println("Aloft API returned an error")
		return
	}

	// Create a new registry for this scrape
	registry := prometheus.NewRegistry()

	// Create labels
	labels := prometheus.Labels{
		"latitude":  latitude,
		"longitude": longitude,
	}

	// Set TFR boolean
	var tfr float64
	if airspace.Data.Overview.Icon == "critical" {
		tfr = 1
	} else {
		tfr = 0
	}

	// Create gauges and set values
	collector := &AirspaceCollector{
		metrics: []*airspaceMetric{
			newAirspaceMetric("airspace_tfr", "Temporary Flight Restriction active", labels, tfr, time.Now()),
		},
	}

	// Register the gauges with the registry
	registry.MustRegister(collector)

	// Use a promhttp.HandlerFor with the new registry to serve the metrics
	promhttp.HandlerFor(registry, promhttp.HandlerOpts{}).ServeHTTP(w, r)
}
