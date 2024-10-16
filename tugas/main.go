package main

import "fmt"

type Kendaraan interface {
	JarakTempuh(bahanBakar float32) float32
}

type Mobil struct {
	Kecepatan float32
}

type Motor struct {
	Kecepatan float32
}

type Bajaj struct {
	Kecepatan float32
}

func (mobil Mobil) JarakTempuh(bbm float32) float32 {
	return bbm * mobil.Kecepatan
}

func (motor Motor) JarakTempuh(bbm float32) float32 {
	return bbm * motor.Kecepatan
}

func (bajaj Bajaj) JarakTempuh(bbm float32) float32 {
	return bbm * bajaj.Kecepatan
}

func terEfisien(bensin float32, kendaraan ...Kendaraan) (float32, Kendaraan) {
	var kendaraanBermotor Kendaraan
	var jarakTerjauh float32 = -1

	for _, wahana := range kendaraan {
		if jarakTempuh := wahana.JarakTempuh(bensin); jarakTerjauh <= jarakTempuh {
			jarakTerjauh = jarakTempuh
			kendaraanBermotor = wahana
		}
	}

	return jarakTerjauh, kendaraanBermotor
}

func main() {
	// contoh 1
	jarakTempuh, kendaraan := terEfisien(10, Mobil{Kecepatan: 1}, Motor{Kecepatan: 3}, Bajaj{Kecepatan: 4})
	fmt.Printf("Kendaraan paling irit adalah %T dengan jarak tempuh: %.2f km\n", kendaraan, jarakTempuh)

	// contoh 2
	jarakTempuh, kendaraan = terEfisien(20, Motor{Kecepatan: 3}, Mobil{Kecepatan: 1})
	fmt.Printf("Kendaraan paling irit adalah %T dengan jarak tempuh: %.2f km\n", kendaraan, jarakTempuh)

	// contoh 3
	jarakTempuh, kendaraan = terEfisien(30, Motor{Kecepatan: 3}, Mobil{Kecepatan: 3}, Bajaj{Kecepatan: 3})
	fmt.Printf("Kendaraan paling irit adalah %T dengan jarak tempuh: %.2f km\n", kendaraan, jarakTempuh)
}
