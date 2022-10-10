package main

import "fmt"

func remDupsSortedArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	output := 1
	last := nums[0]

	for _, curr := range nums[1:] {
		if curr > last {
			nums[output] = curr
			last = curr
			output += 1
		}
	}
	fmt.Println(nums[:output])
	return output
}
