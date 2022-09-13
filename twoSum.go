package main

import "fmt"

/*
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.
You can return the answer in any order.

Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
*/

func twoSumWorst() {
	input := []int{2, 7, 11, 15}
	target := 9

	var result []int
	for i := 0; i < len(input); i++ {
		for j := i + 1; j < len(input); j++ {
			if input[i]+input[j] == target {
				result = append(result, i, j)
			}
		}
	}
	fmt.Println(result)
}

func twoSum(input []int, target int) []int {

	if len(input) == 0 {
		fmt.Println("Input array in empty")
		return nil
	}

	foundMap := make(map[int]int)

	for id, val := range input {
		c := target - val

		if c, ok := foundMap[c]; ok {
			//fmt.Println(input[id], input[c])
			return []int{input[id], input[c]}
		}
		foundMap[val] = id
	}
	return []int{}
}
