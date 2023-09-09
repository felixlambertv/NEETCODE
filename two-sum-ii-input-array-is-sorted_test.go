package neetcode

import (
	"fmt"
	"reflect"
	"testing"
)

func twoSumBruteForce1(numbers []int, target int) []int {
	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j <= len(numbers)-1; j++ {
			if numbers[i]+numbers[j] == target {
				return []int{i + 1, j + 1}
			}
		}
	}
	return nil
}

func TestTwoSumBruteForce1(t *testing.T) {
	testCases := []struct {
		input  []int
		target int
		output []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{1, 2}},
		{[]int{3, 2, 4}, 6, []int{2, 3}},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Input: %v, Target: %d", tc.input, tc.target), func(t *testing.T) {
			res := twoSumBruteForce1(tc.input, tc.target)
			if !reflect.DeepEqual(tc.output, res) {
				t.Errorf("Expected %v, but got %v", tc.output, res)
			}
		})
	}
}

func twoSumBruteOptimize(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left < right {
		fmt.Println(left, right)
		if numbers[left]+numbers[right] > target {
			right--
		} else if numbers[left]+numbers[right] < target {
			left++
		} else {
			return []int{left + 1, right + 1}
		}
	}
	return nil
}

func TestTwoOptimize(t *testing.T) {
	testCases := []struct {
		input  []int
		target int
		output []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{1, 2}},
		{[]int{2, 3, 4}, 6, []int{1, 3}},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Input: %v, Target: %d", tc.input, tc.target), func(t *testing.T) {
			res := twoSumBruteOptimize(tc.input, tc.target)
			if !reflect.DeepEqual(tc.output, res) {
				t.Errorf("Expected %v, but got %v", tc.output, res)
			}
		})
	}
}
