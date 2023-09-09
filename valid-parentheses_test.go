package neetcode

import (
	"testing"
)

func isValid(s string) bool {
	if len(s) < 2 {
		return false
	}
	var stack []byte

	for i := 0; i < len(s); i++ {
		isOpening := s[i] == '{' || s[i] == '[' || s[i] == '('
		if isOpening {
			//push stack
			stack = append(stack, s[i])
		} else if len(stack) == 0 {
			return false
		} else {
			//get last stack value
			prev := stack[len(stack)-1]

			//pop stack
			stack = stack[:len(stack)-1]
			if s[i] == ')' && prev != '(' {
				return false
			} else if s[i] == '}' && prev != '{' {
				return false
			} else if s[i] == ']' && prev != '[' {
				return false
			}
		}
	}

	return len(stack) == 0
}

func TestIsValid(t *testing.T) {
	tests := []struct {
		input  string
		output bool
	}{
		{"()", true},
		{"(())", true},
		{"()[]{}", true},
		{"(]", false},
		{"([)]", false},
		{"{[]}", true},
		{"((", false},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			res := isValid(test.input)
			if res != test.output {
				t.Errorf("Expected %v for input %s, but got %v", test.output, test.input, res)
			}
		})
	}
}
