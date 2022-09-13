package main

import "fmt"

func main() {
	//TwoSum
	//twoSumWorst()

	input := []int{2, 7, 11, 15}
	target := 9
	test := []int{}
	fmt.Println(twoSum(input, target))
	fmt.Println(twoSum(test, target))

}
