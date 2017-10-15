package dataHandling

import (
	"testing"
	"TestProject/config"
	"io/ioutil"
	"os"
)

func TestCreateEntry(t *testing.T) {

	dir := "../" + config.BlogDir
	entryTitle := "TestEntry_Save"
	savedContent := "abcd1234"

	CreateEntry(dir, entryTitle, savedContent)

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



func TestDeleteEntry(t *testing.T) {

	dir := "../" + config.BlogDir
	entryTitle := "TestEntry_Delete"

	err1 := ioutil.WriteFile(dir + entryTitle, []byte(""), 0777)

	if err1 != nil {
		t.Error("Error creating test file: " + err1.Error())
	}

	DeleteEntry(dir, entryTitle)

	if _, err2 := os.Stat(dir + entryTitle); !os.IsNotExist(err2) {
		t.Error("Testfile " + entryTitle + " was not deleted")
	}
}