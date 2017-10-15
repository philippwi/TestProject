package dataHandling

import (
	"io/ioutil"
	"os"
)


func SaveEntry(dir, entry, content string){

	contentConverted := []byte(content)

	err := ioutil.WriteFile(dir + entry, contentConverted, 0777)

	if err != nil {
		panic(err)
	}
}

func DeleteEntry(dir, entry string){

	err := os.Remove(dir + entry)

	if err != nil {
		panic(err)
	}
}