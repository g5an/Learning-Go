package main

import (
	"fmt"
)

func main() {
	defer func() {
		err := recover()
		fmt.Printf("捕捉到了异常%v", err)
	}()
	var a []int
	a[0] = 100 //错误写法
	fmt.Println("这是main函数")
}
