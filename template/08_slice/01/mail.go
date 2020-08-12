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
		"one",
		"two",
		"three",
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", arr)
	if err != nil {
		log.Fatalln(err)
	}
}
