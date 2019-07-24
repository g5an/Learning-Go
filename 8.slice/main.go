package main

import "fmt"

func main() {
	//切片声明
	//1、直接声明
	var a = [3]int{1, 2, 3} //数组
	var b = []int{1, 2, 3}  //切片
	fmt.Println(a)
	fmt.Println(b)
	//2、从数组得到切片
	var c []int
	c = a[1:2]
	fmt.Printf("c:%T\n", c)

	//3、通过make函数构造切片
	e := make([]int, 5, 10) //切片类型    长度   容量
	fmt.Println(e)
	//切片的大小：目前元素的苏良
	fmt.Printf("切片c的大小%d\n", len(c))
	//切片底层数组的容量，cap   //从原始数组切片后的头开始
	fmt.Printf("切片c的容量%d\n", cap(c))

	//切片追加
	/*
		注：切片是引用传递，修改切片内的值，会影响原数组的值
			如果操作的内容超出原数组的容量，切片会在后台复制一个数组进行扩容
			并对新生成的数组操作
			扩容每一次会对元数组容量跟大小x2创建一个新的数组
	*/
	c = append(c, 1, 3, 4)
	fmt.Println(c)
	fmt.Println(a)

	//切片删除元素
	c = append(c[:1], c[2:]...)
	fmt.Println(c)
}
