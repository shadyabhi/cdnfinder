package main

import (
	"sync"
	"testing"
	"time"
)

func TestPrintResponse(t *testing.T) {
	var wg sync.WaitGroup
	duration, _ = time.Duration("3s")

	ns := []nsInfo{
		{
			ip:          "8.8.8.8",
			name:        "Google DNS",
			country_id:  "US",
			city:        "Mountain View",
			reliability: 1.0,
		},
	}
	go printResponse("static.licdn.com", ns, duration, &wg)

	wg.Wait()
}
