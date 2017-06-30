package main

import (
	"reflect"
	"testing"
)

func TestParseCCJSON(t *testing.T) {
	t.Parallel()
	expectedLines := 249

	err := parseCCJSON(ccJSONLoc)
	if err != nil {
		t.Errorf("Error parsing JSON country file: %s", err)
	}
	t.Logf("Successfully parsed JSON country file %s, total items: %d", ccJSONLoc, len(ccInfo))

	linesFound := len(ccInfo)
	if linesFound != expectedLines {
		t.Errorf("Failed to parse JSON country file. Expected lines: %d, found: %d", expectedLines, linesFound)
	}
}

func TestCreateCountryMap(t *testing.T) {
	t.Parallel()
	if err := createCountryMap(); err != nil {
		t.Errorf("Error creating country map: %s", err)
	}
	if _, ok := ccMap["US"]; !ok {
		t.Errorf("Country map was not loaded correctly, US was not found")
	}
	t.Logf("%d keys were loaded", len(reflect.ValueOf(ccMap).MapKeys()))

}
