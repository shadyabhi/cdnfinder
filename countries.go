package main

//Works on this file:-
//	curl https://raw.githubusercontent.com/lukes/ISO-3166-Countries-with-Regional-Codes/master/all/all.json -o countries.json

import (
	"encoding/json"
	"fmt"
)

const ccJSONLoc string = "./countries.json"

type countryCodes []struct {
	Name          string `json:"name"`
	Alpha2        string `json:"alpha-2"`
	Alpha3        string `json:"alpha-3"`
	CountryCode   string `json:"country-code"`
	Iso31662      string `json:"iso_3166-2"`
	Region        string `json:"region"`
	SubRegion     string `json:"sub-region"`
	RegionCode    string `json:"region-code"`
	SubRegionCode string `json:"sub-region-code"`
}

var ccInfo countryCodes
var ccMap map[string]string

func parseCCJSON(loc string) error {
	f, err := readFile(loc)
	if err != nil {
		return fmt.Errorf("Error reading country codes information: %s", err)
	}

	if err := json.Unmarshal(f, &ccInfo); err != nil {
		return fmt.Errorf("Error unmarshalling JSON from file to struct: %s", err)
	}
	return nil
}

func createCountryMap() error {
	err := parseCCJSON(ccJSONLoc)
	if err != nil {
		return fmt.Errorf("Couldn't parse CC json file %s, hence unable to create map: %s", err)
	}
	ccMap = make(map[string]string)

	for _, country := range ccInfo {
		ccMap[country.Alpha2] = country.Name
	}
	return nil
}
