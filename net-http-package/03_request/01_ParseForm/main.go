package main

import (
	"html/template"
	"log"
	"net/http"
)

type demo int

func (d demo) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	tpl.ExecuteTemplate(w, "index.gohtml", req.Form)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var t demo
	http.ListenAndServe(":8080", t)
}
