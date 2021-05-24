package main

import "fmt"

func main() {
	fmt.Println(howManyGames(20, 3, 6, 80))
	fmt.Println(howManyGames(20, 3, 6, 85))
}

func howManyGames(p int32, d int32, m int32, s int32) int32 {
	// Return the number of games you can buy
	count:=int32(0)
	sum:=int32(0)
	if s < p {
		return sum
	}
	for {
		if sum <= s && p <= s-sum {
			sum+=p
			if p-d >= m {
				p-=d
			} else {
				p = m
			}
			count++
		} else {
			return count
		}
	}
}