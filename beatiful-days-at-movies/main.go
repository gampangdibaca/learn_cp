package main

func main() {
	beautifulDays(20, 23, 6)
}

func beautifulDays(i int32, j int32, k int32) int32 {
	// Write your code here
	count:=int32(0)
	for a:=i; a <= j; a++ {
		reversed:=int32(0)
		temp := a
		for {
			if temp <= 0 {
				break
			}
			lastDigit := temp%10
			reversed = reversed * 10 + lastDigit
			temp/=10
		}
		if (a - reversed) % k == 0 {
			count++
		}
	}
	return count
}