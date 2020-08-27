package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":8081", nil)
}

func home(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Method)
}

func about(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Method)
	w.Header().Set("Location", "/")
	w.WriteHeader(http.StatusSeeOther)
}

func contact(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Method)
	tpl, err := template.ParseGlob("template/*")
	if err != nil {
		log.Fatalln(err)
	}

	tpl.ExecuteTemplate(w, "index.html", nil)
}
