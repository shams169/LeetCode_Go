package main

import (
	"sort"
)

/*Problem

Write a function to find the longest common prefix string amongst an array of strings.
If there is no common prefix, return an empty string "".

Example 1:
Input: strs = ["flower","flow","flight"]
Output: "fl"

Example 2:
Input: strs = ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.
*/

func lcp_sorting(strs []string) string {
	sort.Strings(strs)

	first := strs[0]
	last := strs[len(strs)-1]

	res := ""
	for i := 0; i < len(first); i++ {
		if first[i] != last[i] {
			break
		}
		res += string(first[i])
	}
	return res
}
