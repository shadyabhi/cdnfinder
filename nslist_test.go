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

func TestSaveNSFile(t *testing.T) {

	err := saveNSFile(nsURLMock.URL)
	if err != nil {
		t.Errorf("Error saving NS file: %s", err)
	}
}

func TestCreateNSMap(t *testing.T) {
	err := createNSMap(nsURLMock.URL)
	if err != nil {
		t.Errorf("Error creating NS Map: %s", err)
	}
	// US should surely be loaded
	_, ok := nsMap["US"]
	if !ok || len(nsMap["US"]) < 1 {
		t.Errorf("No US nameservers in map, something is definitely wrong")
	}
}

func TestDownloadNS(t *testing.T) {
	// Delete file
	if _, err := os.Stat(nsFile); err == nil {
		err := os.Remove(nsFile)
		if err != nil {
			t.Errorf("Error deleting the %s file: %s", nsFile, err)
		} else {
			t.Logf("Successfully deleted file %s, will execute downloadNS() now", nsFile)
		}
	}

	err := downloadNS(nsURLMock.URL)
	if err != nil {
		t.Errorf("Error downloading Nameserver file: %s", err)
	}
	if _, err := os.Stat(nsFile); os.IsNotExist(err) {
		t.Errorf("downloadNs() was executed but still no file exist")
	}
}
