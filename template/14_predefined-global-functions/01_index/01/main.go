package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	arr := []string{
		"zero",
		"one",
		"two",
		"three",
	}

	err := tpl.Execute(os.Stdout, arr)
	if err != nil {
		log.Fatalln(err)
	}
}
