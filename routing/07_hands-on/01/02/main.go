package main

import (
	"io"
	"log"
	"net/http"
	"text/template"
)

func index(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "index")
}

func dog(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "dog")
}

func me(w http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)

	http.ListenAndServe(":8081", nil)
}
