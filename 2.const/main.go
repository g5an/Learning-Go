package main

import (
	"fmt"
)

const pi = 3.14

// const (
// 	a = 100
// 	b = 1000
// 	c //不写，相当于，等于上一个常量的值 这里c=1000
// 	d //不写，相当于，等于上一个常量的值 这里d=1000
// )

//iota:const中每新增一行常量声明，iota计数+1，
const (
	aa = iota //在const关键字出现的时候重置为0
	bb        //1不写相当于等于上一个值，相当于bb = iota ；而iota每出现一次，iota+1
	cc        //2
	dd = 100
	_  // 变量命名中下划线_代表忽略这个值
	ee //4
)

//iota每新增一行，iota+1
const (
	a, b = iota + 1, iota + 2 //a,b逗号，在同一行，因此这里两个iota都是0
	c, d                      //不写，等于跟上一表达式一样
	e, f
)

func main() {
	//常量不允许修改，修改会报错
	var pi = 2
	fmt.Println(pi)

	fmt.Println(a, b, c, d, e, f, aa, bb, cc, dd, ee)
}

/*
iota ：
1.遇到const,iota初始化为0
2.const中，每新增一行（不是一个变量）声明，iota就递增1
*/
