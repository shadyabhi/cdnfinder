package main

import (
	"testing"
)

func TestReadDb(t *testing.T) {
	t.Parallel()
	f, err := readFile("./data/cnamechain.json")
	if err != nil {
		t.Errorf("Error opening/reading file: %s", err)
	}
	if len(f) != 2459 {
		t.Errorf("File read is malformed")
	}
}

func TestParseJSON(t *testing.T) {
	t.Parallel()
	err := parseCNAMEs("./data/cnamechain.json")
	if err != nil {
		t.Errorf("Error parsing DB file: %s", err)
	}
	if len(cdns.entries) <= 0 {
		t.Errorf("CDN CNAME entries are corrupted")
	}
	t.Logf("CDN CNAME entries loaded: %d", len(cdns.entries))
}

func TestGetCDN(t *testing.T) {
	t.Parallel()
	in := "e9706.dscg.akamaiedge.net"
	out := "Akamai"
	cdn, err := getCDN("e9706.dscg.akamaiedge.net")
	if err != nil {
		t.Errorf("Error getting CDN information from CNAME: %s", err)
	}
	t.Logf("Domain: %s -> CDN: %s", in, cdn)
	if cdn != out {
		t.Errorf("Expected: %s, got: %s", out, cdn)
	}
}
