package main

import (
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/miekg/dns"
)

// getCNAME gets the last CNAME for a record
func getCNAME(domain, ns string) (result string, err error) {
	domain = mkFQDN(domain)

	c := dns.Client{}
	m := dns.Msg{}
	m.SetQuestion(domain, dns.TypeA)
	r, _, err := c.Exchange(&m, ns+":53")
	if err != nil {
		logrus.Errorf("Error contacting DNS Server: %s", err)
		return "", err
	}
	//logrus.Debugf("queryDNS: Took %v for server %s", t, ns)

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
