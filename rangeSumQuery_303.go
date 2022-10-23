package main

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */

type NumArray struct {
	nums []int
}

func Constructor(nums []int) NumArray {
	numarr := NumArray{nums: nums}
	return numarr
}

func (this *NumArray) SumRange(left int, right int) int {
	if left < 0 || right > len(this.nums) {
		return -1
	}
	sum := 0
	for i := left; i <= right; i++ {
		sum += this.nums[i]
	}
	return sum
}
