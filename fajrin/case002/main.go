package main

import "fmt"

type MahlukHidup interface {
	berbicara() string
}

type Orang struct {
	nama string
	umur int
}

func (o Orang) berbicara() string {
	return fmt.Println("Saya Orang Bernama", o.nama, ", Umur Saya", o.umur)
}

type Hewan struct {
	nama string
	umur int
}

func (h Hewan) berbicara() string {
	return fmt.Println("KuruKuru Watashi Wa", h.nama, "Desu, Nenrei Wa", h.umur, "Sai Desu")
}

func speak(m MahlukHidup) string {
	return m.berbicara()
}

func main() {

	fajrin := Orang{"Fajrin", 25}
	kirby := Hewan{"Kirby", 7}

	speak(fajrin)
	speak(kirby)
}
