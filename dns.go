package main

import (
	"strings"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/miekg/dns"
)

// getCNAME gets the last CNAME for a record
func getCNAME(domain, ns string, showDNSErrors bool, dnsTimeout string) (result string, err error) {
	domain = mkFQDN(domain)
	ns = ns + ":53"

	c := dns.Client{}
	// Don't want for more than 2 seconds
	duration, err := time.ParseDuration(dnsTimeout)
	if err != nil {
		logrus.Fatalf("Couldn't parse duration provided for DNS: %s", err)
	}
	c.Timeout = duration

	m := dns.Msg{}
	m.SetQuestion(domain, dns.TypeA)
	r, _, err := c.Exchange(&m, ns)

	if err != nil {
		if showDNSErrors {
			logrus.Errorf("Error contacting DNS Server: %s", err)
		}
		return "", err
	}

	if len(r.Answer) == 0 {
		logrus.Warnf("No results for domain: %s", domain)
	}
	var lastCNAME string
	for _, ans := range r.Answer {
		cname, ok := ans.(*dns.CNAME)
		if ok {
			lastCNAME = cname.Target
		}
	}
	return lastCNAME, nil
}

func mkFQDN(domain string) (result string) {
	result = domain
	if !strings.HasSuffix(domain, ".") {
		result = result + "."
	}
	return result
}
