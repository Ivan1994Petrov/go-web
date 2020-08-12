package main

import (
	"log"
	"os"
	"text/template"
)

type user struct {
	Name   string
	Number string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("tpl.gohtml"))
}

func main() {
	ted := user{
		Name:   "Ted",
		Number: "08888383",
	}

	err := tpl.Execute(os.Stdout, ted)
	if err != nil {
		log.Fatalln(err)
	}
}
