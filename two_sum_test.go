package neetcode

import (
	"fmt"
	"reflect"
	"testing"
)

func twoSumBruteForce(nums []int, target int) []int {
	n := len(nums)
	if n <= 1 {
		return []int{0, 0}
	}

	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if (nums[i] + nums[j]) == target {
				return []int{i, j}
			}
		}
	}
	return []int{0, 0}
}

func TestTwoSumBruteForce(t *testing.T) {
	testCases := []struct {
		nums     []int
		target   int
		expected []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("nums=%v, target=%d", tc.nums, tc.target), func(t *testing.T) {
			res := twoSumBruteForce(tc.nums, tc.target)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, res)
			}
		})
	}
}

func twoSumHashMap(nums []int, target int) []int {
	numsMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		reqNum := target - nums[i]
		idx, isExists := numsMap[reqNum]
		if isExists {
			return []int{idx, i}
		}
		numsMap[nums[i]] = i
	}
	return []int{}
}

func TestTwoSumHashMap(t *testing.T) {
	testCases := []struct {
		nums     []int
		target   int
		expected []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("nums=%v, target=%d", tc.nums, tc.target), func(t *testing.T) {
			res := twoSumHashMap(tc.nums, tc.target)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, res)
			}
		})
	}
}
