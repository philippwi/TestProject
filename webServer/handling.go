package webServer

import (
	"html/template"
	"net/http"
	"TestProject/config"
	"log"
	"fmt"
)

var tpl *template.Template

func StartServer() {
	tpl = template.Must(template.ParseGlob("webServer/*.gohtml"))
	http.HandleFunc("/", idx)
	http.HandleFunc("/apply", aply)
	//http.Handle("/lol", http.NotFoundHandler())
	http.ListenAndServe(config.Port, nil)
}

func idx(w http.ResponseWriter, req *http.Request){
	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	if err != nil{
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func aply(w http.ResponseWriter, req *http.Request){

	if req.Method == http.MethodPost{
		f := req.FormValue("fname")

		fmt.Println(f)
	}
	err := tpl.ExecuteTemplate(w, "apply.gohtml", nil)
	if err != nil{
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}