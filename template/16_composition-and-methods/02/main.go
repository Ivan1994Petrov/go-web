package main

import (
	"log"
	"os"
	"text/template"
)

type user struct {
	Name string
	Age  string
}

type swimmer struct {
	user
	CanSwim bool
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	polly := swimmer{
		user: user{
			Name: "Polly",
			Age:  "26",
		},
		CanSwim: true,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", polly)
	if err != nil {
		log.Fatalln(err)
	}
}
