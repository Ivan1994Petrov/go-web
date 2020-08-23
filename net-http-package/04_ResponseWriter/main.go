package main

import (
	"fmt"
	"net/http"
)

type demo int

func (d demo) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Demo-Key", "this is from demo")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(w, "<h1>Any code you want in this func</h1>")
}

func main() {
	var t demo
	http.ListenAndServe(":8080", t)
}
