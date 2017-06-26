// Source: https://raw.githubusercontent.com/turbobytes/cdnfinder/master/assets/cnamechain.json

package main

import (
	"encoding/json"
	"io/ioutil"
	"strings"
	"sync"
)

// cdnCNAMEs type will hold CDN CNAMEs information
type cdnCNAMEs struct {
	entries [][]string
	sync.Mutex
}

// CDNs holds CDN CNAME information
var cdns cdnCNAMEs

func parseCNAMEs(loc string) error {
	jsonFile, err := readDB(loc)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(jsonFile, &cdns.entries); err != nil {
		return err
	}

	return nil
}

func readDB(loc string) ([]byte, error) {
	raw, err := ioutil.ReadFile(loc)
	if err != nil {
		return raw, err
	}
	return raw, nil
}

func getCDN(domain string) (cdn string, err error) {
	for _, cdnEntry := range cdns.entries {
		if strings.Contains(domain, cdnEntry[0]) {
			cdn = cdnEntry[1]
			break
		}
	}
	return cdn, nil
}
