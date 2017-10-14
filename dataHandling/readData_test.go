package dataHandling

import (
	"testing"
	"TestProject/config"
)

func TestGetEntry(t *testing.T) {
	//Benötigt die Datei TestEntry in /TestProject/BlogEntries/ mit dem Inhalt "Abc123"

	entry := GetEntry("../" + config.BlogDir, "TestEntry")

	if entry != "Abc123"{
		t.Error("Expected: Abc123, got: " + entry)
	}
}