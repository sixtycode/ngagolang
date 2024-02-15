package main

import (
	"log"
	"os"
	"text/template"
)

var tmpl *template.Template

type StrukturData struct {
	Id   int
	Nama string
}

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	data := []StrukturData{
		{1, "Usep"},
		{2, "Sasmita"},
		{3, "Uceng"},
		{4, "Anto"},
		{5, "Rojak"},
		{6, "Saepul"},
		{7, "Roki"},
	}

	err := tmpl.ExecuteTemplate(os.Stdout, "index", data)
	if err != nil {
		log.Fatalln(err)
	}
}
