package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := `第一行
第二行
第三行
	`
	fmt.Println(s1)

	//字符串操作
	s2 := "你好，我是字符串"
	s3 := "hello,i am Hong"

	fmt.Println(len(s2)) //字符串长度
	//字符串拼接
	s4 := fmt.Sprintf("%s--%s", s2, s3)
	fmt.Println(s2 + s3)
	fmt.Println(s4)

	//分割
	s5 := strings.Split(s2, ",")
	fmt.Println(s5)

	//join
	s6 := []string{"wo", "shi", "dsb"}
	fmt.Println(strings.Join(s6, " "))

	//遍历字符串
	s7 := "hello，世界"
	for i := 0; i < len(s7); i++ {
		fmt.Printf("%c\n", s7[i]) //这样输出会把中文变成utf8字符输出
	}
	// 按整个中文字符遍历输出,for range循环是按照rune类型去遍历，自动识别中文

	for k, v := range s7 {
		fmt.Printf("%d%c\n", k, v) //k是索引，可以用下划线_代替k，把索引忽略输出
	}

	// 字符串修改：字符串本身是不可变，要先强制转换成rune或者byte后，一个字节一个字节改才能转

}
