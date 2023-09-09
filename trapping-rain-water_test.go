package neetcode

import (
	"testing"
)

func trap(height []int) int {
	res := 0
	left, right := 0, len(height)-1
	leftMax, rightMax := height[left], height[right]

	for left < right {
		if height[left] < height[right] {
			if height[left] >= leftMax {
				leftMax = height[left]
			} else {
				res += leftMax - height[left]
			}
			left++
		} else {
			if height[right] >= rightMax {
				rightMax = height[right]
			} else {
				res += rightMax - height[right]
			}
			right--
		}
	}

	return res
}

func TestTrap(t *testing.T) {
	tests := []struct {
		input  []int
		output int
	}{
		//{[]int{0, 1, 0, 2, 0, 1, 0}, 2},
		{[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6},
		//{[]int{4, 2, 0, 3, 2, 5}, 9},
		//{[]int{3, 0, 2, 0, 4}, 7},
	}

	for _, test := range tests {
		res := trap(test.input)
		if res != test.output {
			t.Errorf("For input %v, expected %d, but got %d", test.input, test.output, res)
		}
	}
}
