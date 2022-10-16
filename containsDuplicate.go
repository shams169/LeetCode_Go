package main

func containsDuplicate(nums []int) bool {

	contains_dict := make(map[int]int)

	for _, val := range nums {
		if _, ok := contains_dict[val]; ok {
			return true
		}
		contains_dict[val] = 1
	}
	return false
}
