package neetcode

import (
	"reflect"
	"sort"
	"testing"
)

func threeSumOptimize(nums []int) [][]int {
	var res [][]int
	if len(nums) < 3 {
		return res
	}
	sort.Ints(nums)
	for first := 0; first < len(nums)-2; first++ {
		//skip duplicate from left
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}

		second := first + 1
		third := len(nums) - 1

		for second < third {
			sum := nums[first] + nums[second] + nums[third]
			if sum == 0 {
				res = append(res, []int{nums[first], nums[second], nums[third]})
				third--
				//skip duplicate from right
				for second < third && nums[third] == nums[third+1] {
					third--
				}
			} else if sum > 0 {
				third--
			} else {
				second++
			}
		}
	}

	return res
}

func TestThreeSumOptimize(t *testing.T) {
	testCases := []struct {
		input  []int
		output [][]int
	}{
		//{
		//	input:  []int{-1, 0, 1, 2, -1, -4},
		//	output: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		//},
		//{
		//	input:  []int{0, 0, 0, 0},
		//	output: [][]int{{0, 0, 0}},
		//},
		//{
		//	input:  []int{-2, 0, 0, 2, 2},
		//	output: [][]int{{-2, 0, 2}},
		//},
		{
			input:  []int{-1, 0, 1, 2, -1, -4, -2, -3, 3, 0, 4},
			output: [][]int{{-4, 0, 4}, {-4, 1, 3}, {-3, -1, 4}, {-3, 0, 3}, {-3, 1, 2}, {-2, -1, 3}, {-2, 0, 2}, {-1, -1, 2}, {-1, 0, 1}},
		},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			res := threeSumOptimize(tc.input)
			if !reflect.DeepEqual(tc.output, res) {
				t.Errorf("Expected %v, but got %v", tc.output, res)
			}
		})
	}
}

func threeSumBruteForce(nums []int) [][]int {
	var res [][]int
	if len(nums) < 3 {
		return res
	}
	sort.Ints(nums)
	tripletMap := make(map[[3]int][]int)
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				triplet := [3]int{nums[i], nums[j], nums[k]}
				sort.Ints(triplet[:])

				if nums[i]+nums[j]+nums[k] == 0 {
					tripletMap[triplet] = []int{nums[i], nums[j], nums[k]}
				}
			}
		}
	}

	for _, triplet := range tripletMap {
		res = append(res, triplet)
	}

	return res
}

func TestThreeSumBruteForce(t *testing.T) {
	testCases := []struct {
		input  []int
		output [][]int
	}{
		{
			input:  []int{-1, 0, 1, 2, -1, -4},
			output: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			res := threeSumBruteForce(tc.input)
			if !reflect.DeepEqual(tc.output, res) {
				t.Errorf("Expected %v, but got %v", tc.output, res)
			}
		})
	}
}
