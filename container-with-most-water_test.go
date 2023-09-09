package neetcode

import (
	"testing"
)

func maxAreaBruteForce(height []int) int {
	var maxArea int
	left, right := 0, len(height)-1
	for left != right {
		curr := min(height[left], height[right]) * (right - left)
		maxArea = max(maxArea, curr)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maxArea
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func TestMaxAreaBruteForce(t *testing.T) {
	tests := []struct {
		input  []int
		output int
	}{
		{[]int{1, 1}, 1},
		{[]int{1, 2, 3, 4, 5}, 6},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			res := maxAreaBruteForce(test.input)
			if res != test.output {
				t.Errorf("Input: %v, Expected: %d, Got: %d", test.input, test.output, res)
			}
		})
	}
}
