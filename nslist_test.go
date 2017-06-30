package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var nsURLMock *httptest.Server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./testdata/nameservers.csv")
}))
var nsFileTest = "./testdata/nameservers.csv"

func TestSaveNSFile(t *testing.T) {

	err := saveNSFile(nsURLMock.URL)
	if err != nil {
		t.Errorf("Error saving NS file: %s", err)
	}
}

func TestCreateNSMap(t *testing.T) {
	err := createNSMap(nsURLMock.URL, nsFileTest)
	if err != nil {
		t.Errorf("Error creating NS Map: %s", err)
	}
	// US should surely be loaded
	_, ok := nsMap["US"]
	if !ok || len(nsMap["US"]) < 1 {
		t.Errorf("No US nameservers in map, current map: %#v", nsMap)
	}
}

func TestDownloadNS(t *testing.T) {
	err := downloadNS(nsURLMock.URL, nsFileTest)
	if err != nil {
		t.Errorf("Error downloading Nameserver file: %s", err)
	}
	if _, err := os.Stat(nsFile); os.IsNotExist(err) {
		t.Errorf("downloadNs() was executed but still no file exist")
	}
}
