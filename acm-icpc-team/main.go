package main

import "fmt"

func main() {
	fmt.Println(acmTeam([]string{"10101", "11100", "11010", "00101"}))
	fmt.Println(acmTeam([]string{"11101", "10101", "11001", "10111", "10000", "01110"}))
}

func acmTeam(topic []string) []int32 {
	// Write your code here
	totalCount:=int32(0)
	highestCount:=int32(0)
	for i:=0; i<len(topic); i++{
		for j:=i+1; j<len(topic); j++{
			count:=0
			for k:=0; k<len(topic[i]); k++{
				if topic[i][k] == '1' || topic[j][k] == '1' {
					count++
				}
			}
			if highestCount < int32(count) {
				highestCount = int32(count)
				totalCount = 0
			}
			if int32(count) == highestCount {
				totalCount++
			}
		}
	}
	return []int32{highestCount, totalCount}
}