// Source: https://raw.githubusercontent.com/turbobytes/cdnfinder/master/assets/cnamechain.json

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"sync"
)

// cdnCNAMEs type will hold CDN CNAMEs information
type cdnCNAMEs struct {
	entries [][]string
	// initialized is used to find out if entries is laoded or not in an efficient way
	initialized bool
	sync.Mutex
}

// CDNs holds CDN CNAME information
var cdns cdnCNAMEs
var dbFile = "./cnamechain.json"

func parseCNAMEs(loc string) error {
	jsonFile, err := readFile(loc)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(jsonFile, &cdns.entries); err != nil {
		return fmt.Errorf("Error unmarshalling JSON from file to struct: %s", err)
	}
	cdns.initialized = true

	return nil
}

func readFile(loc string) ([]byte, error) {
	raw, err := ioutil.ReadFile(loc)
	if err != nil {
		return raw, err
	}
	return raw, nil
}

func getCDN(domain string) (cdn string, err error) {
	// Initialize cdns if it's not
	if !cdns.initialized {
		err := parseCNAMEs(dbFile)
		if err != nil {
			return "", err
		}
	}
	for _, cdnEntry := range cdns.entries {
		if strings.Contains(domain, cdnEntry[0]) {
			cdn = cdnEntry[1]
			break
		}
	}
	return cdn, nil
}