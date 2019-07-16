package main //声明main包，表明当前是个可执行程序

import (
	"fmt" //导入内置fmt包
	"math"
)

// 在函数外部声明变量，全局可用
var ( //等同于
	name = "Hong" //var name = "Hong"
	age  = 23     //var age = 23
)

func main() { //main函数，是程序执行的入口
	var hig1 = 174                                        //在函数内部声明变量，局部可用
	hig2 := 173                                           //声明短变量简写方式
	fmt.Println("Hello World Println")                    //输出自带换行
	fmt.Print("Hello World Print")                        //输出不带换行
	fmt.Printf("my name is %s,my age is %d\n", name, age) //带有占位符的输出
	fmt.Println(hig1, "and", hig2)

	// 浮点数输出格式
	fmt.Printf("%f\n", math.Pi)   //export 3.141593
	fmt.Printf("%.2f\n", math.Pi) //export 3.14

	// 复数
	var c1 complex64 // complex64 实部和虚部32位
	c1 = 1 + 2i
	var c2 complex128 // complex128 实部和虚部64位
	c2 = 2 + 3i
	fmt.Println(c1)
	fmt.Println(c2)
}
