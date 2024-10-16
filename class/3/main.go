package main

import "fmt"

type Tim struct {
	Nama string
}

func (regu Tim) Cetak() {
	fmt.Println("Nama tim saya adalah", regu.Nama)
}

func main() {
	daftarTim := []Tim{{Nama: "satu"}, {Nama: "dua"}, {Nama: "tiga"}}

	fmt.Println(daftarTim)
	daftarTim[0].Cetak()
}
