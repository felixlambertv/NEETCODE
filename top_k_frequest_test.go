package neetcode

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func topKFrequent(nums []int, k int) []int {
	//Turn to map, put number as key and increment the value
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]] += 1
	}

	//Create new struct so we can sort map
	type kv struct {
		Key int
		Val int
	}
	// Mapping all keys
	var keys []kv
	for key, value := range m {
		keys = append(keys, kv{key, value})
	}
	//sort number count from bigger to smaller
	sort.Slice(keys, func(i, j int) bool {
		return keys[i].Val > keys[j].Val
	})

	var res []int
	for _, key := range keys {
		res = append(res, key.Key)
		if len(res) == k {
			return res
		}
	}

	return []int{}
}

func TestTopKFrequent(t *testing.T) {
	testCases := []struct {
		nums     []int
		k        int
		expected []int
	}{
		{[]int{1, 2}, 2, []int{1, 2}},
		{[]int{1, 1, 1, 2, 2, 3}, 2, []int{1, 2}},
		{[]int{4, 1, 4, 3, 4, 3, 2, 1, 2}, 3, []int{4, 1, 3}},
		{[]int{5, 5, 5, 5, 5, 5, 5, 5, 5, 5}, 1, []int{5}},
		{[]int{1, 2, 3, 4, 5}, 5, []int{1, 2, 3, 4, 5}},
		{[]int{1, 2, 3, 4, 5}, 0, []int{}}, // Edge case: k is 0
		{[]int{}, 3, []int{}},              // Edge case: empty input slice
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("TopKFrequent(%v, %d)", tc.nums, tc.k), func(t *testing.T) {
			res := topKFrequent(tc.nums, tc.k)
			if !reflect.DeepEqual(tc.expected, res) {
				t.Errorf("Expected %v, but got %v", tc.expected, res)
			}
		})
	}
}
