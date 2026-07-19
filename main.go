package main

import (
	"fmt"
	"slices"
)

func findDuplicate(nums []int) int {
	var result int
	slices.Sort(nums)
	i := 0
	j := i + 1

	for {
		if nums[i] != nums[j] {
			i++
			j++
		} else {
			result = nums[i]
			break
		}
	}

	return result
}

func main() {
	nums := []int{1, 3, 4, 2, 2}
	fmt.Println(findDuplicate(nums))
}
