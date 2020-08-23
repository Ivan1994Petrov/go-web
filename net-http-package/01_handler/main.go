package main

import (
	"fmt"
	"net/http"
)

type demo int

func (d demo) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Any code you want in this func")
}

func main() {

}
