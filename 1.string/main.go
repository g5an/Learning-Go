package main

import (
	"fmt"
)

func main() {
	s1 := `第一行
第二行
第三行
	`
	fmt.Println(s1)

	/// 字符串操作
	s2 := "你好，我是字符串"
	s3 := "hello,i am Hong"
	fmt.Println(len(s2))
	s4 := s2 + s3
	fmt.Println(len(s4))
	fmt.Println(s4)

	
}
