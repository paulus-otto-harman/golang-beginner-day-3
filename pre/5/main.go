package main

type Alamat struct {
	Jalan string
	Nomor string
	Kota  string
}

type Rumah struct {
	Alamat
	Warna string
}

func main() {
	satu := Rumah{Alamat: {Jalan: "jalan", Nomor: "1", Kota: ""}, Warna: "Putih"}
}
