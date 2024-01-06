package main

import "fmt"

type identitas struct {
	nama   string
	umur   int
	alamat string
}

func main() {
	data := []identitas{
		{"Ujang", 21, "Jl. Indihiang Raya No. 1 Kota Tasikmalaya"},
		{"Udin", 24, "Jl. Cihideung Asri No. 37 Kota Tasikmalaya"},
		{"Usep", 22, "Jl. Cipatujah Resort No. 7 Kabupaten Tasikmalaya"},
	}

	for _, i := range data {
		fmt.Println("\nHalo nama saya", i.nama, ", umur saya",
			i.umur, "dan saya tinggal di", i.alamat, "\n")
	}
}
