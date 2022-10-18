package main

import "fmt"

/*
Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].
The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.
You must write an algorithm that runs in O(n) time and without using the division operation.
*/

func productExSelf(nums []int) []int {

	inputLen := len(nums)
	prefixArray := make([]int, inputLen)
	postfixArray := make([]int, inputLen)
	result := make([]int, inputLen)

	prefix := 1
	postfix := 1
	for i := 0; i < inputLen; i++ {
		prefixArray[i] = nums[i] * prefix
		prefix = prefixArray[i]
	}
	fmt.Println(prefixArray)

	for i := inputLen - 1; i >= 0; i-- {
		postfixArray[i] = nums[i] * postfix
		postfix = postfixArray[i]
	}

	fmt.Println(postfixArray)
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			result[i] = 1 * postfixArray[i+1]
		} else if i == (inputLen - 1) {
			result[i] = prefixArray[i-1] * 1
		} else {
			result[i] = prefixArray[i-1] * postfixArray[i+1]
		}

	}
	return result
}

func productExSelf_Optimized(nums []int) []int {
	inputLen := len(nums)
	result := make([]int, inputLen)

	prefix := 1
	postfix := 1
	for i := 0; i < inputLen; i++ {
		result[i] = prefix
		prefix *= nums[i]
	}
	//fmt.Println(result)

	for i := inputLen - 1; i >= 0; i-- {
		result[i] *= postfix
		postfix *= nums[i]
	}
	return result
}
