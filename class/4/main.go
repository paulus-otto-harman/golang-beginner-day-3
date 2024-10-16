package main

import "fmt"

func main() {
	satu := 1

	dua := &satu
	tiga := &satu
	fmt.Println(dua, tiga)

	empat := &satu
	fmt.Println("satu dan empat - before :", satu, *empat)
	*empat = 5
	fmt.Println("satu dan empat - after :", satu, *empat)
}
