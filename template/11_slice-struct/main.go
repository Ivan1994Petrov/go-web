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
		Number: "02342",
	}
	po := user{
		Name:   "Po",
		Number: "0989889",
	}
	mo := user{
		Name:   "Mo",
		Number: "333333",
	}

	arr := []user{ted, po, mo}

	err := tpl.Execute(os.Stdout, arr)
	if err != nil {
		log.Fatalln(err)
	}
}
