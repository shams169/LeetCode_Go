package main

import "fmt"

// func findDisappearedNumbers(nums []int) []int {
// 	var result []int
// 	var numsSet = make(map[int]bool)

// 	for i := 0; i < len(nums); i++ {
// 		numsSet[nums[i]] = true
// 	}

// 	for i := 1; i <= len(nums); i++ {
// 		if _, ok := nums[i]; !ok {
// 			result = append(result, i)
// 		}
// 	}
// 	return result
// }

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func findDisappearedNumbers(nums []int) []int {
	var result []int

	for i := 0; i < len(nums); i++ {
		pos := Abs(nums[i]) - 1

		if nums[pos] > 0 {
			nums[pos] *= -1
		}
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			result = append(result, i+1)
		}
	}
	fmt.Println(nums)

	return result
}
