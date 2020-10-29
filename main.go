package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promauto"	
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

var latencyHistogram = promauto.NewHistogram(prometheus.HistogramOpts{
	Name: "httplat_latency_seconds",
})

var interval = time.Duration(10*time.Second)

func measureLatency(url string) {
    ticker := time.NewTicker(interval)
    for _ = range ticker.C {
	    go func() {
		    timer := prometheus.NewTimer(latencyHistogram)
		    defer timer.ObserveDuration()
		    client := &http.Client{Timeout: interval}
		    _, err := client.Get(url)
		    if (err != nil) {
		    	log.Print(err)
		    }
		}()
	}
}

func main() {
	if (len(os.Args) != 2) {
		log.Fatal("Please specify exactly one URL that will be monitored.")
	}
	go measureLatency(os.Args[1])
	http.Handle("/metrics", promhttp.Handler())
	portNum := 9080
	port := os.Getenv("PORT")
	if (port != "") {
		var err error
		portNum, err = strconv.Atoi(port)
		if (err != nil) {
			log.Fatal(err)
		}
	}
	log.Printf("Serving Prometheus metrics on port %d.", portNum)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", portNum), nil))
}
