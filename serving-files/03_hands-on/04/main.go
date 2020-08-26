package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("template/index.html"))
}

func main() {
	http.HandleFunc("/", dogs)
	http.Handle("/res/", http.StripPrefix("/res/", http.FileServer(http.Dir("public"))))

	http.ListenAndServe(":8081", nil)
}

func dogs(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		log.Fatal(err)
	}
}
