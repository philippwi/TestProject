package dataHandling

import (
	"io/ioutil"
)

func GetEntry(dir, entry string) string{
	content, err := ioutil.ReadFile(dir + entry)

	if err != nil {
		panic(err)
		return "Error loading entry: " + err.Error()
	}

	return string(content)
}

func PasswordCorrect (dir, name, enteredPassw string) bool{
	password, err := ioutil.ReadFile(dir + name)

	if err != nil {
		panic(err)
	}

	if enteredPassw == string(password){
		return true
	} else{
		return false
	}
}