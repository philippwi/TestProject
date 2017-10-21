package webServer

import (
	"html/template"
	"net/http"
	"TestProject/config"
	"TestProject/dataHandling"
	"path/filepath"
	"os"
)

var tpl *template.Template

type BlogList struct{
	Blogs []Blog
}

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

	tpl.ExecuteTemplate(w, "home.html", getBlogs())

}

func getBlogs() BlogList{

	blogList := BlogList{}

	filepath.Walk(config.BlogDir, func(path string, f os.FileInfo, err error) error {
		if f.IsDir(){
			return nil
		}
		blogList.Blogs = append(blogList.Blogs, Blog{"",f.Name(),""})
		return nil
	})
	return blogList
}