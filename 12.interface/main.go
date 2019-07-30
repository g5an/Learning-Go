package main

import (
	"fmt"
)

/*
定义接口
*/
type Sayer interface {
	say()
}

type dog struct{}
type cat struct{}

//接口实现   dog跟cat实现了sayer接口中的全部方法，就是接口的实现了
func (d dog) say() {
	fmt.Println("wang! wang! wang! ")
}

func (c cat) say() {
	fmt.Println("miao! miao! miao!")
}

func main() {
	var x Sayer
	a := cat{}
	b := dog{}
	x = a
	x.say()
	x = b
	x.say()
}
