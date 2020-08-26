package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func dog(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "dog")
}

func me(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "me")
}

func main() {
	http.Handle("/", http.HandlerFunc(index))
	http.Handle("/dog/", http.HandlerFunc(dog))
	http.Handle("/me/", http.HandlerFunc(me))

	http.ListenAndServe(":8081", nil)
}
