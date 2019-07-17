package main

import (
	"fmt"
)

func main() {
	s1 := "how do you do"
	s2 := " "
	k := 0
	scoreMap := make(map[string]int, 20)
	for i := 0; i < len(s1); i++ {
		if i == len(s1)-1 {
			s3 := s1[k : i+1]
			k = i + 1
			v, key := scoreMap[s3]
			if key {
				scoreMap[s3] = v + 1
			} else {
				scoreMap[s3] = v + 1
			}
		}
		if s1[i] == s2[0] {
			s3 := s1[k:i]
			k = i + 1
			v, key := scoreMap[s3]
			if key {
				scoreMap[s3] = v + 1
			} else {
				scoreMap[s3] = v + 1
			}
		}
	}
	fmt.Println(scoreMap)
}
