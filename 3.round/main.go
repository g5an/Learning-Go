package main

import (
	"fmt"
)

func main() {
	switchDemo1()
}

func switchDemo1() {
	figer := 3
	switch figer {
	case 1:
		fmt.Println("dmz")
	case 2:
		fmt.Println("sz")
	default:
		fmt.Println("wu")
	}
}
