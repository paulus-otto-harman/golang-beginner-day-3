package main

import "fmt"

type Ayah interface {
	Judol() string
}

type Ibu interface {
	Pinjol() string
}

type Begal struct {
	Nama string
}

func (b *Begal) Judol() string { return "begal taruhan" }

func (b *Begal) Pinjol() string { return "begal ngutang" }

func main() {
	siBegal := Begal{Nama: "x"}

	fmt.Println(siBegal.Nama)
	fmt.Println(siBegal.Judol())
}
