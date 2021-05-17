package main

import "fmt"

func main() {
	fmt.Println(getTotalX([]int32{2, 4}, []int32{16, 32, 96}))
}

func getTotalX(a []int32, b []int32) int32 {
	// Write your code here
	count:=int32(0)
	temp := int32(a[len(a)-1])
	isFactorA := true
	isFactorB := true
	for {
		if temp > b[0] {
			break
		}
		isFactorA = true
		for _, v := range a {
			if temp%v != 0 {
				isFactorA = false
				break
			}
		}
		if !isFactorA {
			temp++
			continue
		}
		isFactorB = true
		for _, v := range b {
			if v%temp != 0{
				isFactorB = false
				break
			}
		}
		if !isFactorB {
			temp++
			continue
		}
		if isFactorA && isFactorB {
			count++
		}
		temp++
	}
	return count
}