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
	ma := map[string]string{
		"ivan":  "petrov",
		"polly": "todorova",
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", ma)
	if err != nil {
		log.Fatalln(err)
	}
}
