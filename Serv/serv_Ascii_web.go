package main

import (
	"html/template"
	"net/http"

	asciiart "../support2"
)

var e *template.Template
var error *template.Template

func Asciiweb(w http.ResponseWriter, req *http.Request) {

	if req.URL.Path != "/" && req.URL.Path != "/ascii-art" {
		errorHandler(w, req, http.StatusNotFound)
		return
	}

	if req.Method != "GET" && req.Method != "POST" {
		errorHandler(w, req, 400)
		return
	}

	err := req.ParseForm()
	if err != nil {
		errorHandler(w, req, 500)
		return
	}

	chaine := req.Form.Get("string")
	banner := req.Form.Get("option")
	resultat := asciiart.AsciiArtWeb2(chaine, banner)

	if resultat == "" {
		errorHandler(w, req, 500)
		return
	}

	e.Execute(w, template.HTML(resultat))
}

func main() {
	e = template.Must(template.ParseFiles("./Ascii_web2.html"))
	error = template.Must(template.New("Error").Parse("{{.}}"))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))
	http.HandleFunc("/", Asciiweb)
	http.HandleFunc("/ascii-art", Asciiweb)
	http.ListenAndServe(":80", nil) //127.0.0.1:80
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		error.Execute(w, template.HTML("Error 404 - Page not found"))
	}
	if status == 400 {
		error.Execute(w, template.HTML("Bad request"))
	}
	if status == 500 {
		error.Execute(w, template.HTML("Internal server error"))
	}
}
