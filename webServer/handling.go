package webServer

import (
	"html/template"
	"net/http"
	"TestProject/config"
	"TestProject/dataHandling"
	"fmt"
	"path/filepath"
	"os"
)

var tpl *template.Template

type Blog struct{
	Author string
	Title string
	Content string
}

func StartServer() {
	tpl = template.Must(template.ParseGlob("webServer/*.html"))
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

	tpl.ExecuteTemplate(w, "index.html", nil)

}

func home(w http.ResponseWriter, req *http.Request){

	if req.Method == http.MethodPost{
		f := req.FormValue("fname")

		fmt.Println(f)
	}

	tpl.ExecuteTemplate(w, "home.html", nil)

}

func getBlogs() []Blog{
	blogs := []Blog{}

	filepath.Walk(config.BlogDir, func(path string, f os.FileInfo, err error) error {
		blogs = append(blogs, Blog{"",path,""})
		return nil
	})


}