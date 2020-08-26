package main

import (
	"io"
	"net/http"
)

func d(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "demo")
}

func t(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "test")
}

func main() {
	http.Handle("/demo", http.HandlerFunc(d))
	http.Handle("/test", http.HandlerFunc(t))

	http.ListenAndServe(":8080", nil)
}
