package dataHandling

import (
	"io/ioutil"
	"time"
	"os"
)

func CreateEntry(dir, entry, content string){

	err := ioutil.WriteFile(dir + entry, []byte(content), 0777)

	if err != nil {
		panic(err)
	}
}

func CreateUser(dir, name, password string){

	if UserExists(dir, name) {
		return
	}

	content :=
		"password: " + password +
		"\ncreationDate: " + time.Now().Format("02.01.2006")

	err := ioutil.WriteFile(dir + name, []byte(content), 0777)

	if err != nil {
		panic(err)
	}
}

func UserExists (dir, name string) bool{
	if _,err := os.Stat(dir+name); os.IsNotExist(err){
		return false
	}

	return true
}

func DeleteEntry(dir, entry string){
	err := os.Remove(dir + entry)

	if err != nil {
		panic(err)
	}
}