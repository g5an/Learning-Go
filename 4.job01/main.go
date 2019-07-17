package main

import "fmt"

//字符串倒序

// s1 = "hello,世界"

// func main(){
// 	i = len(s1)
// 	for k,_ := range s1{
// 		k ,
// 	}
// }

//
func main() {
	var a1 = make([]string, 5, 10)
	var a2 = make([]int, 5, 10)
	var b1 = make([]string, 2, 12)
	var b2 = make([]int, 2, 12)
	fmt.Println(len(a1))
	fmt.Println(a2)
	fmt.Println(b1)
	fmt.Println(b2)

	/*
		输出：     切片string默认空字符    int 默认0
		[    ]
		[0 0 0 0 0]
		[ ]
		[0 0]
	*/
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(a)
	// 	a = append(a, fmt.Sprintf("%v", i))
	// 	b := fmt.Sprintf("%v", i)
	// 	fmt.Println(b)
	// 	fmt.Println(a)
	// }
	// fmt.Println(a)
}
