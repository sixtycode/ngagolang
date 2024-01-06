package main

import "fmt"

type MahlukHidup interface {
	name() string
	age() int
}

type Orang struct {
	nama string
	umur int
}

func (o Orang) name() string {
	return o.nama
}

func (o Orang) age() int {
	return o.umur
}

func (h Hewan) name() string {
	return h.nama
	// fmt.Println("Saya Orang Bernama", o.nama)
	// , ", Umur Saya", o.umur)
}

func (h Hewan) age() int {
	return h.umur
}

type Hewan struct {
	nama string
	umur int
}

// func (h Hewan) berbicara() string {
// 	return fmt.Println("KuruKuru Watashi Wa", h.nama, "Desu, Nenrei Wa", h.umur, "Sai Desu")
// }

func speak(m MahlukHidup) {
	fmt.Println("Saya Orang Bernama", m.name(), ", Umur Saya", m.age())
}

func main() {

	fajrin := Orang{"Fajrin", 25}
	kirby := Hewan{"Kirby", 7}

	speak(fajrin)
	speak(kirby)
}
