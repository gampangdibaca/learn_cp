package main

import (
	"fmt"
)

func main() {
	a := "test"
	fmt.Println(a[:3] + "wow")
	fmt.Println(repeatedString("ojowrdcpavatfacuunxycyrmpbkvaxyrsgquwehhurnicgicmrpmgegftjszgvsgqavcrvdtsxlkxjpqtlnkjuyraknwxmnthfpt", 685118368975))
	//fmt.Println(repeatedString("ababa", 3))
}

func repeatedString(s string, n int64) int64 {
	// Write your code here
	count := int64(0)
	aInString := int64(0)
	for _, v := range s {
		if v == 'a' {
			aInString++
		}
	}
	if n > int64(len(s)) {
		count = (n / int64(len(s))) * aInString
		excess := n%int64(len(s))
		if excess != 0 {
			for _, v := range s[:excess] {
				if v == 'a' {
					count++
				}
			}
		}
		return count
	}
	for _, v := range s[:n] {
		if v == 'a' {
			count++
		}
	}
	return count
}
