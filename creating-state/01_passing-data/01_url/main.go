package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	v := req.FormValue("q")

	fmt.Fprintln(w, "Do my search: "+v)
}
