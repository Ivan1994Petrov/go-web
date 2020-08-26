package main

import (
	"io"
	"net/http"
)

type test int

func (t test) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/dog":
		io.WriteString(w, "Dog")
	case "/cat":
		io.WriteString(w, "Cat")
	}
}

func main() {
	var t test

	http.ListenAndServe(":8080", t)
}
