package serverHandling

import (
	"fmt"
	"net/http"
	"TestProject/config"
)




func StartServer() {

	fmt.Println("http://localhost" + config.Port)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	http.ListenAndServe(config.Port, nil)
}