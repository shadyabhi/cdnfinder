package main

import (
	"testing"
)

func TestReadDb(t *testing.T) {
	f, err := readDB("./cnamechain.json")
	if err != nil {
		t.Errorf("Error opening/reading file: %s", err)
	}
	if len(f) != 2459 {
		t.Errorf("File read is malformed")
	}
}

func TestParseJSON(t *testing.T) {
	err := parseCNAMEs("./cnamechain.json")
	if err != nil {
		t.Errorf("Error parsing DB file: %s", err)
	}
	if len(cdns.entries) <= 0 {
		t.Errorf("CDN CNAME entries are corrupted")
	}
	t.Logf("CDN CNAME entries loaded: %d", len(cdns.entries))
}
