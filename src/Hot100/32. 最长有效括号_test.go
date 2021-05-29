package Hot100

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type LongestValidParenthesesInput struct {
	s string
}

type LongestValidParenthesesOutput struct {
	ret int
}

func TestLongestValidParentheses(t *testing.T) {
	inputs := []LongestValidParenthesesInput{
		{"())"},
		{"("},
		{"))))())()()(()"},
		{"()(()"},
		{"(()"},
		{")()())"},
		{""},
	}

	expected := []LongestValidParenthesesOutput{
		{2},
		{0},
		{4},
		{2},
		{2},
		{4},
		{0},
	}

	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		ret := longestValidParentheses(input.s)
		assert.Equal(t, expected[i].ret, ret)
	}

}
