package main

import (
	"io"
	"net/http"
)

type test int

func (t test) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "test")
}

type demo int

func (d demo) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "demo")
}

func main() {
	var t test
	var d demo

	mux := http.NewServeMux()
	mux.Handle("/test/", t)
	mux.Handle("/demo", d)

	http.ListenAndServe(":8080", mux)
}
