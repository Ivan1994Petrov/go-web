package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/set", set)
	http.HandleFunc("/read", read)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":8081", nil)
}

func home(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("my-cookie")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(c)
	fmt.Fprintln(w, "Cookie:", c)
}

func set(w http.ResponseWriter, req *http.Request) {
	c := &http.Cookie{
		Name:  "my-cookie",
		Value: "some value",
	}
	http.SetCookie(w, c)
	fmt.Println(c)
	fmt.Fprintln(w, "YOUR COOKIE:", c)
	fmt.Fprintln(w, "COOKIE WRITTEN - CHECK YOUR BROWSER")
	fmt.Fprintln(w, "in chrome go to: dev tools / application / cookies")
}

func read(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("my-cookie")
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	fmt.Println(http.StatusBadRequest)
	fmt.Fprintln(w, "YOUR COOKIE:", c)
}
