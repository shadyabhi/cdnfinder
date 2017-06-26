package main

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/olekukonko/tablewriter"
)

// printResponse prints CDN info to stdout, no error is checked
func query(domain string, nameservers []nsInfo, showDNSErrors bool, dnsTimeout time.Duration, results chan Results, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	for _, ns := range nameservers {
		cname, took, err := getCNAME(domain, ns.ip, showDNSErrors, dnsTimeout)
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
		if ok && cdn != "" {
			// logrus.Infof("%s (%s:%s): %s (%s)", countryName, ns.country_id, ns.city, cdn, cname)
			result := Results{
				cdn:          cdn,
				countryName:  countryName,
				nameserver:   ns,
				cname:        cname,
				responseTime: took,
			}
			results <- result
		}
	}
}

func printTable(results chan Results, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()

	table := tablewriter.NewWriter(os.Stdout)
	headers := []string{"Country", "City", "CDN/Hostname", "IP", "Time"}
	table.SetColWidth(60)
	table.SetHeader(headers)
	table.SetBorders(tablewriter.Border{Left: true, Top: true, Right: true, Bottom: true})
	table.SetCenterSeparator("|")

	for r := range results {
		city := fmt.Sprintf("%s - %-15s", r.nameserver.country_id, r.nameserver.city)
		cdnHostname := fmt.Sprintf("%-10s - %s", r.cdn, r.cname)
		table.Append([]string{r.countryName, city, cdnHostname, r.nameserver.ip, r.responseTime.String()})
	}

	table.Render()
}
