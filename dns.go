package main

import (
	"strings"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/miekg/dns"
)

// getCNAME gets the last CNAME for a record
func getCNAME(domain, ns string, showDNSErrors bool, dnsTimeout time.Duration) (result string, took time.Duration, err error) {
	domain = mkFQDN(domain)
	ns = ns + ":53"

	c := dns.Client{}
	// Don't want for more than 2 seconds
	c.Timeout = dnsTimeout

	m := dns.Msg{}
	m.SetQuestion(domain, dns.TypeA)
	r, took, err := c.Exchange(&m, ns)

	if err != nil {
		if showDNSErrors {
			logrus.Errorf("Error contacting DNS Server: %s", err)
		}
		return "", 0, err
	}

	if len(r.Answer) == 0 {
		logrus.Warnf("No results for domain: %s", domain)
		return "", 0, err
	}
	var lastCNAME string
	for _, ans := range r.Answer {
		cname, ok := ans.(*dns.CNAME)
		if ok {
			lastCNAME = cname.Target
		}
	}
	return lastCNAME, took, nil
}

func mkFQDN(domain string) (result string) {
	result = domain
	if !strings.HasSuffix(domain, ".") {
		result = result + "."
	}
	return result
}
