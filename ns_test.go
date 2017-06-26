package main

import (
	"os"
	"testing"
)

func TestSaveNSFile(t *testing.T) {
	err := saveNSFile()
	if err != nil {
		t.Errorf("Error saving NS file: %s", err)
	}
}

func TestCreateNSMap(t *testing.T) {
	err := createNSMap()
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

	err := downloadNS()
	if err != nil {
		t.Errorf("Error downloading Nameserver file: %s", err)
	}
	if _, err := os.Stat(nsFile); os.IsNotExist(err) {
		t.Errorf("downloadNs() was executed but still no file exist")
	}
}
