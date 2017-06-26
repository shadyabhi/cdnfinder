package main

import (
	"flag"
	"sync"
	"time"

	"github.com/Sirupsen/logrus"
)

func main() {
	// Commandline parameters
	showDNSErrors := flag.Bool("errors", false, "Show DNS errors in output?")
	domain := flag.String("domain", "", "Domain name to investigate")
	dnsTimeout := flag.String("timeout", "3s", "Seconds to waut for DNS responses")
	flag.Parse()

	// Check sanity of options
	if *domain == "" {
		logrus.Errorf("No '-domain' specified, exiting")
		flag.Usage()
		logrus.Fatalf("")
	}

	logrus.Infof("Starting query at time: %s", time.Now())

	// Create NS map
	err := createNSMap()
	if err != nil {
		logrus.Fatalf("Couldn't create map of NS: %s", err)
	}

	// Create Country Map
	err = createCountryMap()
	if err != nil {
		logrus.Fatalf("Couldn't create map of Countrues: %s", err)
	}

	var wg sync.WaitGroup

	for _, nameservers := range nsMap {
		// use only first ns
		go printResponse(*domain, nameservers[:1], *showDNSErrors, *dnsTimeout, &wg)
	}

	// Wait for all goroutines to print
	wg.Wait()
}
