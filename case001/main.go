package main

import (
	"log"
	"os"
	"text/template"
)

var tmpl *template.Template

type StrukturData struct {
	Nama string
	Umur int
}

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	data := []StrukturData{
		{"Udin", 19},
		{"Ujang", 24},
		{"Usro", 31},
	}

	err := tmpl.ExecuteTemplate(os.Stdout, "index", data)
	if err != nil {
		log.Fatalln(err)
	}
}
