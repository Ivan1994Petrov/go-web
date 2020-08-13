package main

import (
	"log"
	"os"
	"text/template"
)

type user struct {
	Name string
	Age  int
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	ivan := user{
		Name: "Ivan",
		Age:  26,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", ivan)
	if err != nil {
		log.Fatalln(err)
	}

}
