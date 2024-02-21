package main

import (
	"log"
	"os"
	"text/template"
)

var tmpl *template.Template

type StrukturData struct {
	Id    int
	Nama  string
	Nilai float32
}

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	data := []StrukturData{
		{1, "Usep Karawita", 73.7},
		{2, "Sasmita Sumbul", 81.4},
		{3, "Uceng Kawazaki", 83.6},
		{4, "Anto Estriver", 44.1},
		{5, "Rojak Kago", 62.9},
		{6, "Saepul Khidir", 91.6},
		{7, "Roki Khasmiri", 84.9},
		{8, "Ismail Kanabawi", 66.6},
		{9, "El Gato", 94.4},
		{10, "Thomas Katt", 53.8},
		{11, "Antent Zaslavski", 88.3},
		{12, "Senn Cross", 80.0},
		{13, "Alvin Nathanael", 69.2},
		{14, "Samantha D. Larsson", 71.5},
	}

	err := tmpl.ExecuteTemplate(os.Stdout, "index", data)
	if err != nil {
		log.Fatalln(err)
	}
}
