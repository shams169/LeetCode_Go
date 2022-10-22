package main

import (
	"fmt"
	"sort"
)

func missingNumber(nums []int) int {
	length := len(nums)

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	fmt.Println(nums)
	if nums[0] != 0 {
		return 0
	}

	if nums[length-1] != length {
		return length
	}

	for i := 0; i < length; i++ {
		if nums[i+1]-nums[i] != 1 {
			return nums[i] + 1
		}
	}

	return -1
}
