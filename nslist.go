package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/Sirupsen/logrus"
)

const nsURL string = "https://public-dns.info/nameservers.csv"
const nsFile string = "./data/nameservers.csv"

type nsInfo struct {
	ip          string
	name        string
	country_id  string
	city        string
	reliability float64
}

var nsMap map[string][]nsInfo

func saveNSFile(nsURL string) (err error) {
	start := time.Now()

	resp, err := http.Get(nsURL)
	if err != nil {
		return fmt.Errorf("Error issuing GET request to fetch NS: %s", err)
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			logrus.Warningf("Error closing body, don't care: %s", err)
		}
	}()

	file, err := os.Create(nsFile)
	if err != nil {
		return fmt.Errorf("Error creating file to save nameservers list %s", err)
	}
	n, err := io.Copy(file, resp.Body)
	if err != nil {
		return fmt.Errorf("Error writing nameservers to file from response: %s", err)
	}
	elasped := time.Since(start)

	logrus.Debugf("Successfully wrote file %s (%d bytes) in %s", nsFile, n, elasped)
	return nil
}

func createNSMap() error {
	// Download Nameserver file if needed
	err := downloadNS(nsURL)
	if err != nil {
		logrus.Fatalf("Unable to fetch %s or it doesn't exist, can't proceed further without that: %s", nsFile, err)
	}

	file, err := os.Open(nsFile)
	if err != nil {
		return fmt.Errorf("Failed while opening NS file: %s", err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			logrus.Warningf("Error closing file, might be corrupted: %s", err)
		}
	}()

	nsMap = make(map[string][]nsInfo)
	reader := csv.NewReader(file)

	// read header of CSV as it contains column names
	if _, err := reader.Read(); err != nil {
		logrus.Fatalf("Couldn't even read header of CSV file, something is really wrong: %s", err)
	}

	// Create map
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			// Couldn't parse a line? Don't care
			thisErr, ok := err.(*csv.ParseError)
			if !ok {
				logrus.Warningf("Got a non csv.ParseError error while parsing json")
				continue
			}
			logrus.Warningf("Skipping a line while parsing csv at line: %d, column: %d", thisErr.Line, thisErr.Column)
		}
		if len(record) < 10 {
			// Bad server sends us incomplete CSVs
			// Generally the last line is always incomplete
			continue
		}

		//logrus.Debugf("Current CSV line to parse: %s", record)

		countryID := record[2]
		reliability, err := strconv.ParseFloat(record[7], 64)
		if err != nil {
			// We won't add this record
			logrus.Warningf("Won't add record, unable to parse reliability to float. Record was: %s", record[7])
			continue
		}

		ns := nsInfo{
			ip:          record[0],
			name:        record[1],
			country_id:  countryID,
			city:        record[3],
			reliability: reliability,
		}
		nsMap[countryID] = append(nsMap[countryID], ns)
	}
	return nil
}

// downloadNS downloads NS file if it doesn't exist
func downloadNS(nsURL string) error {
	if _, err := os.Stat(nsFile); os.IsNotExist(err) {
		logrus.Warningf("%s file was not found, downloading...", nsFile)
		err := saveNSFile(nsURL)
		if err != nil {
			return fmt.Errorf("Couldn't fetch NS file from https://public-dns.info/")
		}
	}
	return nil
}
