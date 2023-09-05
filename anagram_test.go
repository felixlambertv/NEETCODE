package neetcode

import (
	"sort"
	"strings"
	"testing"
)

func isAnagram(s string, t string) bool {
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

func TestAnagram(t *testing.T) {
	s1 := "anagram"
	s2 := "nagaram"
	res := isAnagram(s1, s2)
	if res != true {
		t.Fail()
	}
}
