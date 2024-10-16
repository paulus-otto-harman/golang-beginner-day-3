package main

import "fmt"

type User struct {
	Nama   string
	Alamat string
	Telpon string
}

func inisialisasi(nama string, alamat string, telpon string) User {
	return User{Nama: nama, Alamat: alamat, Telpon: telpon}
}

func main() {
	saya := inisialisasi("Amir", "Menteng", "081")
	fmt.Println(saya)
}
