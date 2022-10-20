package main

/*
Given an integer array nums, find the contiguous subarray (containing at least one number)
which has the largest sum and return its sum.

A subarray is a contiguous part of an array.

Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6.
*/

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxSubArray(input []int) int {

	maxSum := input[0]
	currentSum := maxSum

	for i := 1; i < len(input); i++ {
		currentSum = maxInt(input[i], (input[i] + currentSum))
		maxSum = maxInt(currentSum, maxSum)
	}
	return maxSum
}
