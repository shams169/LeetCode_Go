package main

func singleNumber(nums []int) int {
	numsMap := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if _, ok := numsMap[nums[i]]; !ok {
			numsMap[nums[i]] = 1
		} else {
			numsMap[nums[i]] += 1
		}
	}

	for k, v := range numsMap {
		if v == 1 {
			return k
		}
	}

	return 0
}
