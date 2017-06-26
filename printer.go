package main

import (
	"sync"

	"github.com/Sirupsen/logrus"
)

// printResponse prints CDN info to stdout, no error is checked
func printResponse(domain string, nameservers []nsInfo, showDNSErrors bool, dnsTimeout string, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	for _, ns := range nameservers {
		cname, err := getCNAME(domain, ns.ip, showDNSErrors, dnsTimeout)
		if err != nil {
			if showDNSErrors {
				logrus.Warningf("Couldn't resolve domain: %s via ns: %s, Error: %s", domain, ns, err)
			}
			continue
		}

		cdn, err := getCDN(cname)
		if err != nil {
			logrus.Warningf("Couldn't get CDN name from hostname %s, Error: %s", cname, err)
		}
		countryName, ok := ccMap[ns.country_id]
		// Only show complete information
		if ok {
			logrus.Infof("%s (%s:%s): %s (%s)", countryName, ns.country_id, ns.city, cdn, cname)
		}
	}
}
