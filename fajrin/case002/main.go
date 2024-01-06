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
	// fmt.Println("Saya Orang Bernama", o.nama)
	// , ", Umur Saya", o.umur)
}

func (o Orang) age() int {
	return o.umur
}

// type Hewan struct {
// 	nama string
// 	umur int
// }

// func (h Hewan) berbicara() string {
// 	return fmt.Println("KuruKuru Watashi Wa", h.nama, "Desu, Nenrei Wa", h.umur, "Sai Desu")
// }

func speak(m MahlukHidup) {
	fmt.Println(m.name())
	fmt.Println(m.age())
}

func main() {

	// fajrin := Orang{"Fajrin", 25}
	// kirby := Hewan{"Kirby", 7}

	var fajrin Orang
	fajrin.nama = "Fajrin"
	fajrin.umur = 25

	speak(fajrin)

	// speak(fajrin)
	// speak(kirby)
}
