package main

import "fmt"

type mahlukHidup interface {
	berbicara() string
}

type orang struct {
	nama string
	umur int
}

func (o orang) berbicara() string {
	fmt.Println("Saya Orang Bernama", o.nama, ", Umur Saya", o.umur)
}

type hewan struct {
	nama string
	umur int
}

func (h hewan) berbicara() string {
	fmt.Println("KuruKuru Watashi Wa", h.nama, "Desu, Nenrei Wa", h.umur, "Sai Desu")
}


func speak(m mahlukHidup) berbicara() string{
	fmt.Println(m.)
}

func main() {
	var 
	fmt.Println("Halo Dunia")
	fajrin.berbicara()
	kirby.berbicara()

	speak()

}
 