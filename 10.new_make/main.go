package main

import (
	"fmt"
)

//new 是用来初始化，值类型指针的
//make是用来初始化slice、map、channels
func main() {
	var a = new(int)
	fmt.Println(a)
	fmt.Println(*a)

	var c = new([3]int)
	fmt.Println(c)
	c[0] = 1
	fmt.Println(*c)
}
