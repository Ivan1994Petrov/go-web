package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

type saga struct {
	Name  string
	Motto string
}

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	if len(s) >= 3 {
		s = s[:3]
	}
	return s
}

func main() {
	b := saga{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
	}
	g := saga{
		Name:  "Gandhi",
		Motto: "Be the change",
	}
	sagas := []saga{b, g}

	data := struct {
		Wisdom []saga
	}{
		sagas,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}
}
