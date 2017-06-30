package main

import (
	"flag"
	"strings"
	"sync"
	"time"

	"github.com/Sirupsen/logrus"
)

type Results struct {
	cdn          string
	countryName  string
	cname        string
	nameserver   nsInfo
	responseTime time.Duration
}

func main() {
	start := time.Now()

	// Setup logging
	logrus.SetFormatter(&logrus.TextFormatter{})
	logrus.SetLevel(logrus.DebugLevel)

	// Commandline parameters
	showDNSErrors := flag.Bool("errors", false, "Show DNS errors in output?")
	domain := flag.String("domain", "", "Domain name to investigate")
	timeout := flag.String("timeout", "3s", "Seconds to waut for DNS responses")
	numberNS := flag.Int("number-ns", 1, "Number of nameservers to query from each country")
	filterCN := flag.String("filter-countries", "", "Comma-separated list of countries to include. Default behaviour is empty and shows all countries")
	flag.Parse()

	// Check sanity of options
	if *domain == "" {
		logrus.Errorf("No '-domain' specified, exiting")
		flag.Usage()
		logrus.Fatalf("")
	}
	if *numberNS < 0 {
		logrus.Errorf("Illegal value provided for '-number-ns'")
		flag.Usage()
		logrus.Fatalf("")
	}
	dnsTimeout, err := time.ParseDuration(*timeout)
	if err != nil {
		logrus.Errorf("Cannot parse specified '-timeout'")
	}
	includeCNs := strings.Split(*filterCN, ",")
	// End of Commandline parameters

	// SETUP before we start execution
	// Parse CDN CNAMEs
	err = parseCNAMEs(dbFile)
	if err != nil {
		logrus.Fatalf("Couldn't parse CDN CNAMEs file: %s", err)
	}

	// Create NS map
	err = createNSMap(nsURL)
	if err != nil {
		logrus.Fatalf("Couldn't create map of NS: %s", err)
	}

	// Create Country Map
	err = createCountryMap()
	if err != nil {
		logrus.Fatalf("Couldn't create map of Countrues: %s", err)
	}
	// END of SETUP

	// results channel is sent data from `query` function and
	// is received at `printTable` to print a table
	results := make(chan Results, 10000)
	var wgPrint sync.WaitGroup
	wgPrint.Add(1)
	go printTable(results, &wgPrint)

	var wgQuery sync.WaitGroup

	totalNS := 0
	for cn, nameservers := range nsMap {
		if *filterCN != "" && !sliceContains(includeCNs, cn) {
			continue
		}
		nameserversLength := len(nameservers)

		// Only query the number that's available, no panic!
		var nsToQuery int
		if *numberNS > nameserversLength {
			nsToQuery = nameserversLength
		} else {
			nsToQuery = *numberNS
		}
		totalNS = totalNS + nsToQuery
		// Each ns in it's own goroutine
		for _, ns := range nameservers[:nsToQuery] {
			wgQuery.Add(1)
			go query(*domain, ns, *showDNSErrors, dnsTimeout, results, &wgQuery)
		}
	}

	wgQuery.Wait()
	// Close channel once all queries are done
	close(results)

	wgPrint.Wait()

	elasped := time.Since(start)
	logrus.Infof("Took %s to query %d DNS servers...", elasped, totalNS)
}
