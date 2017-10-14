package serverHandling

import (
	"fmt"
	"net/http"
)

const Port = ":80"

func PrintLink(){
	x := "5"
	Port += x
}

func StartServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	http.ListenAndServe(Port, nil)
}