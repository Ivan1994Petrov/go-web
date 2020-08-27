package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/shome", secondHome)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":8081", nil)
}

func home(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Method)
}

func secondHome(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Method)
	http.Redirect(w, req, "/", http.StatusMovedPermanently)
}
