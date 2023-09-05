package neetcode

import (
	"fmt"
	"reflect"
	"testing"
)

func productExceptSelfBruteForce(nums []int) []int {
	res := make([]int, 0, len(nums))
	for i := 0; i < len(nums); i++ {
		count := 1
		for j := 0; j < len(nums); j++ {
			if i != j {
				count *= nums[j]
			}
		}
		res = append(res, count)
	}

	return res
}

func productExceptSelf(nums []int) []int {
	size := len(nums)
	left := make([]int, size)
	right := make([]int, size)
	answer := make([]int, size)
	left[0] = 1
	right[size-1] = 1
	for i := 1; i < size; i++ {
		j := size - i - 1
		left[i] = nums[i-1] * left[i-1]
		right[j] = nums[j+1] * right[j+1]
	}
	fmt.Println(left, right)
	for i := 0; i < size; i++ {
		answer[i] = left[i] * right[i]
	}
	return answer
}

func TestProductExceptSelf(t *testing.T) {
	tests := []struct {
		input  []int
		output []int
	}{
		{[]int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
		//{[]int{-1, 1, 0, -3, 3}, []int{0, 0, 9, 0, 0}},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Input: %v", test.input), func(t *testing.T) {
			res := productExceptSelf(test.input)
			if !reflect.DeepEqual(test.output, res) {
				t.Errorf("Expected %v but got %v", test.output, res)
			}
		})
	}
}
