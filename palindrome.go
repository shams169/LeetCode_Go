package main

import (
	"strconv"
)

func isPalindrome(x int) bool {

	//Convert the int to string
	num := strconv.FormatInt(int64(x), 10)

	i := 0
	j := len(num) - 1

	for i < j {
		if num[i] != num[j] {
			return false
		}
		i += 1
		j -= 1

	}

	return true
}
