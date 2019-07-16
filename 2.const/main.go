package main

import (
	"fmt"
)

const pi = 3.14
const (
	a = 100
	b = 1000
	c //不写，相当于，等于上一个常量的值 这里c=1000
	d //不写，相当于，等于上一个常量的值 这里d=1000
)

//iota:const中每新增一行常量声明，iota计数+1，
const (
	aa = iota //在const关键字出现的时候重置为0
	bb        //1不写相当于等于上一个值，相当于bb = iota ；而iota每出现一次，iota+1
	cc        //2
	dd = 100
	_  // 变量命名中下划线_代表忽略这个值
	ee //4
)

func main() {
	//常量不允许修改，修改会报错
	fmt.Println(pi)

	fmt.Println(a, b, c, d, aa, bb, cc, dd, ee)
}
