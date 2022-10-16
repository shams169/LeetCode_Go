package main

import "fmt"

func main() {
	//TwoSum
	//twoSumWorst()

	/*
		input := []int{2, 7, 11, 15}
		target := 9
		test := []int{}
		fmt.Println(twoSum(input, target))
		fmt.Println(twoSum(test, target))
	*/

	// Check Palindrome
	/*
		fmt.Println(isPalindrome(121))
		fmt.Println(isPalindrome(132231))
		fmt.Println(isPalindrome(1009))
		fmt.Println(isPalindrome(-109))
		fmt.Println(isPalindrome(1))
	*/

	// Roman to Integer
	/*
		fmt.Println(RomanToInt("III"))
		fmt.Println(RomanToInt("IV"))
		fmt.Println(RomanToInt("LVIII"))
		fmt.Println(RomanToInt("V"))
		fmt.Println(RomanToInt("VIII"))
	*/

	// LongestCommonPrefix

	// fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	// fmt.Println(longestCommonPrefix([]string{"fxwer", "flow", "flowmerht"}))

	// fmt.Println(LCP([]string{"flower", "flow", "flight"}))
	// fmt.Println(LCP([]string{"fxwer", "flow", "flowmerht"}))

	// fmt.Println(lcprefix([]string{"flower", "flow", "flight"}))

	// fmt.Println(lcp_sorting([]string{"flower", "flow", "flight"}))
	// fmt.Println(lcp_sorting([]string{"dog", "race", "flight"}))

	//************************************************//
	// Valid Parantheses
	//************************************************//

	// fmt.Println(validParans("{()}"))
	// fmt.Println(validParans("[]"))
	// fmt.Println(validParans("[)]"))
	// fmt.Println(validParans("}()[]}"))

	//************************************************//
	// Remove Duplicates from Sorted Array
	//************************************************//

	//fmt.Println(remDupsSortedArray([]int{1, 2, 2, 3, 3, 3, 4, 4, 5, 6, 6}))

	//************************************************//
	// Remove Duplicates from Sorted Array
	//************************************************//

	//removeElement([]int{1, 2, 2, 3, 3, 3, 4, 4, 5, 6, 2}, 2)

	//************************************************//
	// Buy Sell Stock at Max profit
	//************************************************//

	// fmt.Println(buySellStock_Brute([]int{7, 1, 5, 3, 6, 4}))
	// fmt.Println(buySellStock_Brute([]int{7, 6, 4, 3, 1}))

	fmt.Println(buySellStock_Optimized([]int{1, 7, 5, 3, 6, 4}))
	fmt.Println(buySellStock_Optimized([]int{7, 6, 4, 3, 1}))
}
