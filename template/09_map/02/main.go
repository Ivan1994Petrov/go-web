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
	mp := map[string]int{
		"ivan":  11111,
		"polly": 2222,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", mp)
	if err != nil {
		log.Fatalln(err)
	}
}
