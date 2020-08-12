package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("tpl.gohtml"))
}

func main() {
	slice := []string{
		"zero",
		"one",
		"two",
		"three",
	}

	data := struct {
		Words []string
		Lname string
	}{
		slice,
		"Ted",
	}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}

}
