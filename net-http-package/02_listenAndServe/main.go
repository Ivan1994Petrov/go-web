package main

import (
	"fmt"
	"net/http"
)

type demo int

func (d demo) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Println(w, "Any code you want")
}

func main() {
	var t demo
	http.ListenAndServe(":8080", t)
}
