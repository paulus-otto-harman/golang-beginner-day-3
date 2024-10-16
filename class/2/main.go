package main

import "fmt"

type Tim struct {
	Nama string
	Asal string
}

func main() {
	timSurabaya := Tim{Nama: "Menteng", Asal: "Jakarta Pusat"}
	fmt.Println(timSurabaya)
}
