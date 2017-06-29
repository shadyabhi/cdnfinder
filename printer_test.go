package main

import (
	"sync"
	"testing"
	"time"
)

func TestPrintResponse(t *testing.T) {
	var wg sync.WaitGroup
	duration, _ := time.ParseDuration("3s")

	ns := nsInfo{
		ip:          "8.8.8.8",
		name:        "Google DNS",
		country_id:  "US",
		city:        "Mountain View",
		reliability: 1.0,
	}

	results := make(chan Results)
	go query("static.licdn.com", ns, false, duration, results, &wg)
	wg.Wait()

	result := <-results
	t.Logf("Queries static.licdn.com, got %s", result)
}
