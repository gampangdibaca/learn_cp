package main

import "fmt"

func main() {
	fmt.Println(birthday([]int32{1, 2, 1, 3, 2}, 3, 2))
	fmt.Println(birthday([]int32{1, 1, 1, 1, 1, 1}, 3, 2))
	fmt.Println(birthday([]int32{4}, 4, 1))
}

func birthday(s []int32, d int32, m int32) int32 {
	count:=int32(0)
	for i, _ := range s{
		if int32(i) + m - 1 < int32(len(s)) {
			temp:=int32(0)
			for j := int32(i); j < int32(i) + m; j++ {
				temp += s[j]
			}
			if temp == d {
				count++
			}
		}
	}
	return count
}
