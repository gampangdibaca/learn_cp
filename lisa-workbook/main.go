package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(workbook(5 , 3, []int32{4, 2, 6, 1, 10}))
	fmt.Println(workbook(10 , 5, []int32{3, 8, 15, 11, 14, 1, 9, 2, 24, 31}))
}

func workbook(n int32, k int32, arr []int32) int32 {
	// Write your code here
	currPage:=int32(0)
	lclp:=int32(0)
	count:=int32(0)
	for _, v := range arr {
		cp:= int32(math.Ceil(float64(v)/float64(k)))
		for i:=int32(0);i<cp;i++{
			currPage++
			if i==cp-1{
				temp:=v%k
				if temp==0{
					temp=k
				}
				if currPage <= v && currPage >= v-temp+1 {
					count++
					fmt.Printf("%d >= %d <= %d\n", v, currPage, v-temp+1)
				}
			} else if (currPage-lclp)*k>=currPage && (currPage-lclp)*k-k+1<=currPage{
				count++
				fmt.Printf("%d >= %d <= %d\n", (currPage-lclp)*k, currPage, (currPage-lclp)*k-k+1)
			}
		}
		lclp=currPage
	}
	return count
}