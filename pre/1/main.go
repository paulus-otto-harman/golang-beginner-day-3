package main

import "fmt"

func main() {
	a := 1
	b := a
	fmt.Println("----", b)

	b = 2
	fmt.Println(b)
	fmt.Println(a)

	c := &a
	fmt.Println("----", *c)

	*c = 2
	fmt.Println(*c)
	fmt.Println(a)
}
