package main

func main() {

}

func minimumNumber(n int32, password string) int32 {
	// Return the minimum number of characters to make the password strong
	var count int32
	if len(password) < 6 {
		count = int32(6) - int32(len(password))
	} else {
		count = 0
	}
	isStrong := []bool{false, false, false, false}
	for _, v := range password {
		if !isStrong[0] && int(v) >= 97 && int(v) <= 122 {
			isStrong[0] = true
		} else if !isStrong[1] && int(v) >= 65 && int(v) <= 90 {
			isStrong[1] = true
		} else if !isStrong[2] && int(v) >= 48 && int(v) <= 57 {
			isStrong[2] = true
		} else if !isStrong[3] && (int(v) <= 45 || int(v) == 64 || int(v) == 94) {
			isStrong[3] = true
		}
	}
	temp:=int32(0)
	for _, v := range isStrong {
		if !v {
			temp++
		}
	}
	if len(password) > 6 {
		count+=temp
	} else {
		if temp > count {
			count = temp
		}
	}
	return count
}
