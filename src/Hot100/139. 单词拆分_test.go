package Hot100

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type WordBreakInput struct {
	s        string
	wordDict []string
}

type WordBreakOutput struct {
	ret bool
}

func TestWordBreak(t *testing.T) {
	inputs := []WordBreakInput{
		{"ccbb", []string{"bc", "cb"}},
		{"leetcode", []string{"leet", "code"}},
		{"applepenapple", []string{"apple", "pen"}},
		{"catsandog	", []string{"cats", "dog", "sand", "and", "cat"}},
	}

	expected := []WordBreakOutput{
		{false},
		{true},
		{true},
		{false},
	}

	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		ret := wordBreak(input.s, input.wordDict)
		assert.Equal(t, expected[i].ret, ret)
	}
}
