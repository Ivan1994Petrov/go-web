package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "index")
}

func dog(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "dog")
}

func me(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "me")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/god/", dog)
	http.HandleFunc("/me/", me)

	http.ListenAndServe(":8081", nil)
}
