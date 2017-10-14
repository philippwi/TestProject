package dataHandling

import (
	"io/ioutil"
)


func SaveEntry(dir, entry, content string) bool{

	contentConverted := []byte(content)

	err := ioutil.WriteFile(dir + entry, contentConverted, 0664)

	if err != nil {
		panic(err)
		return false
	}

	return true
}
