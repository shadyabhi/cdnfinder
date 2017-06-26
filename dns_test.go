package main

import "testing"

func TestMkFQDN(t *testing.T) {
	var ttCases = []struct {
		in  string
		out string
	}{
		{"www.google.com", "www.google.com."},
		{"www.google.com.", "www.google.com."},
	}
	for _, tt := range ttCases {
		if tt.out != mkFQDN(tt.in) {
			t.Errorf("%s != %s", tt.in, tt.out)
		}
	}
}

func TestQueryDNS(t *testing.T) {
	in := "torrent.abhi.host"
	out := "abhijeetr.com."

	cname, err := getCNAME(in, "8.8.8.8")
	if err != nil {
		t.Errorf("Error getting CNAME record: %s", err)
	}
	if cname != out {
		t.Errorf("CNAME is not expected. Expected %s, found %s", out, cname)
	}
}
