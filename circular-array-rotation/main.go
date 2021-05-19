package main

import "fmt"

func main() {
	a := []int32 {39356, 87674, 16667, 54260, 43958, 70429, 53682, 6169, 87496, 66190, 90286, 4912, 44792, 65142, 36183,
		43856, 77633, 38902, 1407, 88185, 80399, 72940, 97555, 23941, 96271, 49288, 27021, 32032, 75662, 69161, 33581,
		15017, 56835, 66599, 69277, 17144, 37027, 39310, 23312, 24523, 5499, 13597, 45786, 66642, 95090, 98320, 26849,
		72722, 37221, 28255, 60906}
	fmt.Println(circularArrayRotation(a, 51, []int32{47, 10, 12, 13, 6}))
	fmt.Println(circularArrayRotation([]int32{1, 2, 3, 4}, 1, []int32{0, 1, 2, 3}))
	fmt.Println(circularArrayRotation([]int32{1, 2, 3, 4}, 2, []int32{0, 1, 2, 3}))
	fmt.Println(circularArrayRotation([]int32{1, 2, 3, 4}, 3, []int32{0, 1, 2, 3}))
	fmt.Println(circularArrayRotation([]int32{1, 2, 3, 4}, 4, []int32{0, 1, 2, 3}))
	fmt.Println(circularArrayRotation([]int32{1, 2, 3, 4}, 5, []int32{0, 1, 2, 3}))
	fmt.Println(circularArrayRotation([]int32{1, 2, 3}, 1, []int32{0, 1, 2}))

}

func circularArrayRotation(a []int32, k int32, queries []int32) []int32 {
	// Write your code here
	var newArray []int32
	if int32(len(a)) == k || k%int32(len(a)) == 0{
		newArray = a
	} else {
		if k > int32(len(a)){
			k = k % int32(len(a))
		}
		k = int32(len(a)) - k
		newArray = a[k:]
		newArray = append(newArray, a[:k]...)
	}

	queried := []int32{}
	for _, v := range queries {
		queried = append(queried, newArray[v])
	}
	return queried
}
