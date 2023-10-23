package neetcode

import (
	"fmt"
	"testing"
)

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func TestSearch(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		expect int
	}{
		{
			nums:   []int{-1, 0, 3, 5, 9, 12},
			target: 9,
			expect: 4,
		},
		{
			nums:   []int{-1, 0, 3, 5, 9, 12},
			target: 2,
			expect: -1,
		},
	}

	for _, test := range tests {
		result := search(test.nums, test.target)
		if result == test.expect {
			fmt.Printf("PASSED: For input nums = %+v and target = %d, got: %d\n", test.nums, test.target, result)
		} else {
			fmt.Printf("FAILED: For input nums = %+v and target = %d, expected: %d but got: %d\n", test.nums, test.target, test.expect, result)
		}
	}
}
