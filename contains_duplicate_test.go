package neetcode

import (
	"sort"
	"testing"
)

func containDuplicateBruteForce(nums []int) bool {
	n := len(nums)
	if n <= 1 {
		return false
	}

	for i := 0; i < n-1; i++ {
		counter := i + 1
		for j := counter; j < n; j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}
	return false
}

func TestBruteForce(t *testing.T) {
	testInput := []int{1, 2, 3, 4, 5, 5, 6, 7, 8, 9}
	res := containDuplicateBruteForce(testInput)
	if res != true {
		t.Fail()
	}
}

func containDuplicateSortFirst(nums []int) bool {
	n := len(nums)
	if n <= 1 {
		return false
	}
	sort.Ints(nums)
	for i := 0; i < n-1; i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}
	return false
}

func TestSortFirst(t *testing.T) {
	testInput := []int{1, 2, 3, 4, 5, 5, 6, 7, 8, 9}
	res := containDuplicateSortFirst(testInput)
	if res != true {
		t.Fail()
	}
}
