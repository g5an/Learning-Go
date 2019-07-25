package main

import (
	"fmt"
)

/*
函数格式：
func 函数名 (参数) (返回值){    //多个返回值用逗号,分开，
	函数体
}
*/
func intSum(x int, y int) int {
	return x + y
}

/*
可变参数: 参数后面加... 一般放最后，用来接收剩余的所有参数
多返回值： 多返回值用用括号括起来，逗号分来
*/
func intsum2(x int, y int, z ...int) (int, int) {
	fmt.Printf("z type is %T\n", z) //是一个切片[]int
	sum := 0
	for _, v := range z {
		sum = sum + v
	}
	return sum, sum + x + y
}

func localvar() {
	var num = 40
	fmt.Println("localvar num is :", num)
}

func main() {
	//函数调用
	a1 := 6
	b1 := 8
	c1 := intSum(a1, b1)
	d1 := intSum(8, 9)
	fmt.Println(c1, d1)

	//可变参数调用
	a2, b2 := intsum2(10, 20, 30, 20, 40)
	fmt.Println(a2, b2)

	//闭包调用
	f := funname("shahe")
	f("xiao")
}

//闭包:内层函数调用外层函数的变量
func funname(varname1 string) func(returnname string) {
	return func(returnname string) {
		fmt.Printf("hello,%v", returnname)
		fmt.Println(varname1)
	}
}
