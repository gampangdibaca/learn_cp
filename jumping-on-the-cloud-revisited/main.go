package main

import "fmt"

func main() {
	fmt.Println(jumpingOnClouds([]int32{0, 0, 1, 0, 0, 1, 1, 0}, 2))
	fmt.Println(jumpingOnClouds([]int32{1, 1, 1, 0, 1, 1, 0, 0, 0, 0}, 3))

}

func jumpingOnClouds(c []int32, k int32) int32 {
	energy:=int32(100)
	for i:=int32(0); ; {
		if c[(i+k)%int32(len(c))] == 1 {
			energy -= 3
		} else {
			energy--
		}
		i = (i+k)%int32(len(c))
		if i == 0 {
			return energy
		}
	}
}