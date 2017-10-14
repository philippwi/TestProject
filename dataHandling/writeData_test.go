package dataHandling

import (
	"testing"
	"TestProject/config"
	"io/ioutil"
	"os"
)

func TestSaveEntry(t *testing.T) {

	dir := "../" + config.BlogDir
	entryTitle := "TestEntry_Save"
	savedContent := "abcd1234"

	SaveEntry(dir, entryTitle, savedContent)

	readContent, err := ioutil.ReadFile(dir + entryTitle)

	if err != nil {
		t.Error("Error reading file: " + err.Error())
	}

	if string(readContent) != savedContent{
		t.Error("Expected: " + savedContent + ", got: " + string(readContent))
	}

	err = os.Remove(dir + entryTitle)

	if err != nil{
		t.Error("Function test passed but error deleting test file: " + err.Error())
	}
}