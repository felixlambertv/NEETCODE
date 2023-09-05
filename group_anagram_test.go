package neetcode

import (
	"reflect"
	"sort"
	"testing"
)

//func groupAnagrams(strs []string) [][]string {
//	if len(strs) == 1 {
//		return [][]string{}
//	}
//
//	groupedString := groupBasedOnLength(strs)
//	for _, words := range groupedString {
//		for i := 0; i < len(words); i++ {
//			for j := 0; j < len(words[i]); j++ {
//				ss := strings.Split(words[i], "")
//				sort.Strings(ss)
//			}
//		}
//	}
//	fmt.Println(groupedString)
//	return [][]string{}
//}

func groupAnagrams(strs []string) [][]string {
	ag := make(map[string][]string)
	for _, str := range strs {
		sortedStr := sortString(str)
		ag[sortedStr] = append(ag[sortedStr], str)
	}

	var res [][]string
	for _, a := range ag {
		res = append(res, a)
	}
	return res
}

func sortString(s string) string {
	sBytes := []byte(s)
	sort.Slice(sBytes, func(i, j int) bool {
		return sBytes[i] < sBytes[j]
	})
	return string(sBytes)
}

func TestGroupAnagrams(t *testing.T) {
	testCases := []struct {
		input  []string
		output [][]string
	}{
		{
			input:  []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			output: [][]string{{"bat"}, {"tan", "nat"}, {"ate", "eat", "tea"}},
		},
		{
			input:  []string{"listen", "silent", "hello", "world"},
			output: [][]string{{"listen", "silent"}, {"hello"}, {"world"}},
		},
		{
			input:  []string{"abc", "def", "xyz"},
			output: [][]string{{"abc"}, {"def"}, {"xyz"}},
		},
	}

	for _, testCase := range testCases {
		res := groupAnagrams(testCase.input)
		if !reflect.DeepEqual(res, testCase.output) {
			t.Errorf("Input: %v, Expected: %v, Got: %v", testCase.input, testCase.output, res)
		}
	}
}
