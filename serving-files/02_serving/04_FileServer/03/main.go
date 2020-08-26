package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", dog)
	http.Handle("/res/", http.StripPrefix("/res/", http.FileServer(http.Dir("./assets"))))
	http.ListenAndServe(":8081", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	io.WriteString(w, `<img src="/res/toby.jpg">`)
}
