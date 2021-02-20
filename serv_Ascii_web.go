package main

import (
	"html/template"
	"net/http"
)

type page struct {
	Module string
}

var e *template.Template

func Asciiweb(w http.ResponseWriter, req *http.Request) {
	e.Execute(w, page{
		Module: "Ascii_Art_WEB",
	})
}

func main() {
	e = template.Must(template.ParseFiles("./Ascii_web.html"))
	http.HandleFunc("/", Asciiweb)
	http.ListenAndServe(":80", nil)
}