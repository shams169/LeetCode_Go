package main

import "fmt"

func removeElement(nums []int, val int) {

	//{1, 2, 2, 3, 3, 3, 4, 4, 5, 6, 2}
	i := 0
	j := 0

	for j < len(nums) {
		if nums[j] == val {
			j += 1
		} else {
			nums[i] = nums[j]

			i += 1
			j += 1
		}
	}

	fmt.Println(nums[:i])

}
