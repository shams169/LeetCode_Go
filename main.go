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

	fmt.Println(lcp_sorting([]string{"flower", "flow", "flight"}))
	fmt.Println(lcp_sorting([]string{"dog", "race", "flight"}))

}

// func LCP(strs []string) string {

// 	minLen := len(strs[0])

// 	// for _, val := range strs {
// 	// 	if len(val) < minLen {
// 	// 		minLen = len(val)
// 	// 	}
// 	// }

// 	lcp := ""
// 	var found bool = true
// 	for i := 0; i < minLen; i++ {
// 		currentChar := string(strs[0][i])

// 		for _, val := range strs {
// 			if string(val[i]) != currentChar {
// 				found = false
// 				break
// 			}
// 		}
// 		if !found {
// 			break
// 		} else {
// 			lcp += currentChar
// 		}
// 	}
// 	return lcp

// }
