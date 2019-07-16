package main

import "fmt"

//无法直接定义二进制，要通过10禁止转换
var (
	a = 10   //十进制
	b = 077  //八进制  数字零开头
	c = 0xff //十六进制 0x开头
)

func main() {

	fmt.Println(a, b, c)  // int 默认打印十进制
	fmt.Printf("%b\n", a) //无法直接定义二进制，要通过10进制转换
	fmt.Printf("%o\n", b)
	fmt.Printf("%x\n", c)

}
