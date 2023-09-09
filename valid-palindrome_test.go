package neetcode

import (
	"strings"
	"testing"
	"unicode"
)

func isPalindromeStartFromMiddle(s string) bool {
	cleanString := strings.Map(func(r rune) rune {
		if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
			return -1
		}
		return unicode.ToLower(r)
	}, s)
	mid := len(cleanString) % 2
	left := ((len(cleanString) - mid) / 2) - 1
	right := (len(cleanString) + mid) / 2

	for i := 1; i <= len(cleanString)/2; i++ {
		if cleanString[left] != cleanString[right] {
			return false
		}
		left--
		right++
	}

	return true
}

func isPalindrome(s string) bool {
	cleanString := strings.Map(func(r rune) rune {
		if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
			return -1
		}
		return unicode.ToLower(r)
	}, s)
	left, right := 0, len(cleanString)-1
	for left < right {
		if cleanString[left] != cleanString[right] {
			return false
		}
		left++
		right--
	}

	return true
}

func TestIsPalindrome(t *testing.T) {
	testCases := []struct {
		input  string
		output bool
	}{
		{"A man, a plan, a canal: Panama", true},
		{"racecar", true},
		{"hello", false},
		{"Was it a car or a cat I saw?", true},
		{"No 'x' in Nixon", true},
		{"12321", true},
		{"ab", false},
		{"0P", false},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			res := isPalindrome(tc.input)
			if res != tc.output {
				t.Errorf("Expected %v, but got %v", tc.output, res)
			}
		})
	}
}
