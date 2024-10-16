package main

import "fmt"

type People struct {
	Komunitas []Person
}

type Person struct {
	Tahun int
}

func (p *People) add(person Person) {
	p.Komunitas = append(p.Komunitas, person)
}

func filter(people People, tahun int) []Person {
	list := []Person{}
	for _, person := range people.Komunitas {
		if person.Tahun == tahun {
			list = append(list, person)
		}
	}

	return list
}

func main() {
	daftar := People{}
	daftar.add(Person{Tahun: 2020})
	daftar.add(Person{Tahun: 2021})
	daftar.add(Person{Tahun: 2022})
	fmt.Println(daftar)

	fmt.Println(filter(daftar, 2020))
}
