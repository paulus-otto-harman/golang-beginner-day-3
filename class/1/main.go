package main

import "fmt"

type Pertandingan struct {
	Nama string
	Skor string
}

func (pertandingan *Pertandingan) gantiSkor(skor string) {
	pertandingan.Skor = skor
}

func main() {
	laga := Pertandingan{Nama: "Bogor vs Depok", Skor: "1-0"}
	laga.gantiSkor("1-3")

	fmt.Println(laga)
}
