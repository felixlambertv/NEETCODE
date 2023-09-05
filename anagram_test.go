package neetcode

import (
	"sort"
	"strings"
	"testing"
)

func isAnagram1(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	n := len(s)
	ss := strings.Split(s, "")
	ts := strings.Split(t, "")
	sort.Strings(ss)
	sort.Strings(ts)
	for i := 0; i < n; i++ {
		if ss[i] != ts[i] {
			return false
		}
	}
	return true
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	r1 := []rune(s)
	sort.Slice(r1, func(i, j int) bool {
		return r1[i] < r1[j]
	})

	t1 := []rune(t)
	sort.Slice(t1, func(i, j int) bool {
		return t1[i] < t1[j]
	})

	if string(r1) == string(t1) {
		return true
	}
	return false
}

func TestAnagram(t *testing.T) {
	s1 := "anagram"
	s2 := "nagaram"
	res := isAnagram(s1, s2)
	if res != true {
		t.Fail()
	}
}
