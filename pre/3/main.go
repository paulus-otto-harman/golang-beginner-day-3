package main

import "fmt"

type Ayah interface {
	Hobi() string
	Pekerjaan() string
}

type Ibu interface {
	Hobi() string
	Kegiatan() string
}

type Keturunan interface {
	Ayah
	Ibu
}

type Anak struct {
	Nama string
}

func (a Anak) Hobi() string {
	return "main game"
}

func (a Anak) Pekerjaan() string {
	return "pelajar"
}

func (a Anak) Kegiatan() string {
	return "nonton tv"
}

func hobiPekerjaanDanKegiatan(k Keturunan) {
	fmt.Println(k.Hobi())
	fmt.Println(k.Pekerjaan())
	fmt.Println(k.Kegiatan())
}

func main() {
	siAnak := Anak{Nama: "Dodi"}
	hobiPekerjaanDanKegiatan(siAnak)

	//fmt.Println(siAnak.Hobi())

	//pekerjaanDanKegiatan(siAnak)
}
