package Hot100

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type LengthOfLongestSubstringInput struct {
	str string
}

type LengthOfLongestSubstringOutput struct {
	ret int
}

func TestLengthOfLongestSubstring(t *testing.T) {
	inputs := []LengthOfLongestSubstringInput{
		{"tmmzuxt"},
		{"adebcbfghi"},
		{"abcabcbb"},
		{"bbbbb"},
		{"pwwkew"},
		{""}}

	expected := []LengthOfLongestSubstringOutput{
		{5},
		{6},
		{3},
		{1},
		{3},
		{0}}

	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		ret := lengthOfLongestSubstring(input.str)
		assert.Equal(t, expected[i].ret, ret)
	}

}
