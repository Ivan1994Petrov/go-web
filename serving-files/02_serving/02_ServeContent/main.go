package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", dog)
	http.HandleFunc("/toby.jpg", dogPic)

	http.ListenAndServe(":8081", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	io.WriteString(w, `<img src="/toby.jpg">`)
}

func dogPic(w http.ResponseWriter, req *http.Request) {
	file, err := os.Open("toby.jpg")
	if err != nil {
		http.Error(w, "missing file", 404)
		return
	}

	defer file.Close()

	file1, err := file.Stat()
	if err != nil {
		http.Error(w, "missing file", 404)
		return
	}

	http.ServeContent(w, req, file.Name(), file1.ModTime(), file)
}
