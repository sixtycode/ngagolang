package main

import "fmt"

type orang struct {
	nama string
	umur int
}

type hewan struct {
	nama string
	umur int
}

type mahlukHidup interface {
}

func berbicara() {

}

func speak(m mahlukHidup) {

}

func main() {
	fmt.Println("Halo Dunia")

}
