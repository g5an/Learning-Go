package main

import "fmt"

func main() {
	//声明时要指定 大小，跟变量类型
	var a = [3]string{"hello", "world", "and"} //声明时初始化
	var b [3]string                            //声明时不初始化，string默认填充空字符，int默认填充0
	var c = [...]int{1, 2, 3, 4, 5, 6}         //声明时候指定内容，不指定数组大小，由编译器去判断数组大小
	b = [3]string{"hel", "lo", "world"}

	//根据索引值初始化
	var d [100]int
	d = [100]int{10: 1} //第100位填1，其它按默认值0 填
	d[1] = 2
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
