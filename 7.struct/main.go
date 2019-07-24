package main

import (
	"fmt"
)

type person struct {
	name string
	city string
	age  int8
}

func main() {

	var p1 = new(person)
	p1.name = "hong"
	fmt.Printf("%T\n", p1)
	fmt.Printf("%#v\n", p1)
}
