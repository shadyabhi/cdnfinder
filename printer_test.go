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

	results := make(chan Results, 10)
	// Start consumer, otherwise it will block
	go func() {
		result := <-results
		t.Logf("Queries static.licdn.com, got %s", result)
	}()

	// Producer
	wg.Add(1)
	go func() {

		err := query("static.licdn.com", ns, false, duration, results, &wg)
		if err != nil {
			t.Errorf("Woah, we have an error")
		}
	}()
	wg.Wait()

	close(results)
}
