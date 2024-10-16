package main

import "fmt"

type BangunDatar interface {
	Keliling() float32
}

//1 function 1 parameter
//function dapat menerima 3 object bangun datar
//function berfungsi menghitung keliling bangun datar

type PersegiPanjang struct {
	Panjang float32
	Lebar   float32
}

type Lingkaran struct {
	JariJari float32
}

type Segitiga struct {
	SisiAb float32
	SisiBc float32
	SisiCa float32
}

func (persegiPanjang PersegiPanjang) Keliling() float32 {
	return 2 * (persegiPanjang.Lebar + persegiPanjang.Lebar)
}

func (lingkaran Lingkaran) Keliling() float32 {
	const pi = 3.14
	return 2 * pi * lingkaran.JariJari
}

func (segitiga Segitiga) Keliling() float32 {
	return segitiga.SisiAb + segitiga.SisiBc + segitiga.SisiCa
}

func keliling(bangunDatar BangunDatar) float32 {
	return bangunDatar.Keliling()
}

func main() {
	pp := PersegiPanjang{Panjang: 5, Lebar: 10}
	l := Lingkaran{JariJari: 10}
	st := Segitiga{SisiAb: 5, SisiBc: 10, SisiCa: 10}

	fmt.Println("Keliling Persegi Panjang", keliling(pp))
	fmt.Println("Keliling Lingkaran", keliling(l))
	fmt.Println("Keliling Segitiga", keliling(st))
}
