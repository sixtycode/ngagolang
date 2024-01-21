package main

import "fmt"

type MahlukHidup interface {
	berbicara()
}

type Orang struct {
	nama string
	umur int
}

type Hewan struct {
	nama string
	umur int
}

func (o Orang) berbicara() {
	fmt.Println("Saya Orang Bernama", o.nama, ", Umur Saya", o.umur)
}

func (h Hewan) berbicara() {
	fmt.Println("KuruKuru Watashi Wa", h.nama, "Desu, Nenrei Wa", h.umur, "Sai Desu")
}

func speak(m MahlukHidup) {
	m.berbicara()
}

func main() {
	fajrin := Orang{"Fajrin", 25}
	kirby := Hewan{"Kirby", 7}
	speak(fajrin)
	speak(kirby)
}
