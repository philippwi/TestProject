package webServer

import (
	"html/template"
	"net/http"
	"TestProject/config"
	"TestProject/dataHandling"
	"log"
	"fmt"
)

var tpl *template.Template

func StartServer() {
	tpl = template.Must(template.ParseGlob("webServer/*.gohtml"))
	http.HandleFunc("/", login)
	http.HandleFunc("/home", home)
	//http.Handle("/lol", http.NotFoundHandler())
	http.ListenAndServe(config.Port, nil)
}

func login(w http.ResponseWriter, r *http.Request){

	if r.Method == http.MethodPost{
		userName := r.FormValue("fuser")
		password := r.FormValue("fpass")

		if !dataHandling.UserExists(config.UserDir, userName){
			dataHandling.CreateUser(config.UserDir,userName,password)
			http.Redirect(w, r, "/home", 302)
		}else{
			if dataHandling.PasswordCorrect(config.UserDir, userName,password){
				http.Redirect(w, r, "/home", 302)
			}else{
				http.Redirect(w, r, "/", 302)
			}
		}
	}

	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	if err != nil{
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func home(w http.ResponseWriter, req *http.Request){

	if req.Method == http.MethodPost{
		f := req.FormValue("fname")

		fmt.Println(f)
	}
	err := tpl.ExecuteTemplate(w, "home.gohtml", nil)
	if err != nil{
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}