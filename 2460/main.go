package main

import (
	"fmt"
)

func applyOperations(nums []int) []int {
	for i := range nums {
		if i+1 == len(nums) {
			break
		}
		if nums[i] == nums[i+1] {
			nums[i] = nums[i] * 2
			nums[i+1] = 0
		}
	}
	result := []int{}
	numberOfZeros := 0
	for j := range nums {
		if nums[j] == 0 {
			numberOfZeros += 1
		} else {
			result = append(result, nums[j])
		}
	}
	for range numberOfZeros {
		result = append(result, 0)
	}

	return result
}

func main() {
	example := []int{847, 847, 0, 0, 0, 399, 416, 416, 879, 879, 206, 206, 206, 272}

	fmt.Println(applyOperations(example))
}
