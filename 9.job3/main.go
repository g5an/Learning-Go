package main

import (
	"fmt"
)

/*
你有50枚金币，需要分配给以下几个人：Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth。
分配规则如下：
a. 名字中每包含1个'e'或'E'分1枚金币
b. 名字中每包含1个'i'或'I'分2枚金币
c. 名字中每包含1个'o'或'O'分3枚金币
d: 名字中每包含1个'u'或'U'分4枚金币
写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
程序结构如下，请实现 ‘dispatchCoin’ 函数
*/
var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
	name         = [...]string{"e", "E", "i", "I", "o", "O", "u", "U"}
)

func dispatchCoin() int {
	sum := 0
	for _, k := range users {
		distribution[k] = 0
		for _, key := range k {
			key := string(rune(key))
			v, _ := distribution[k]
			switch key {
			case name[0], name[1]:
				distribution[k] = v + 1
				sum = sum + 1
			case name[2], name[3]:
				distribution[k] = v + 2
				sum = sum + 2
			case name[4], name[5]:
				distribution[k] = v + 3
				sum = sum + 3
			case name[6], name[7]:
				distribution[k] = v + 4
				sum = sum + 4
			default:
				continue
			}
		}
	}
	return coins - sum
}
func main() {
	left := dispatchCoin()
	fmt.Println("剩下：", left)
	fmt.Println(distribution)
}
