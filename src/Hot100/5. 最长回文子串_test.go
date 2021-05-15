package Hot100

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type LongestPalindromeInput struct {
	str string
}

type LongestPalindromeOutput struct {
	ret string
}

func TestLongestPalindrome(t *testing.T) {
	inputs := []LongestPalindromeInput{
		{"aaaa"},
		{"babad"},
		{"cbbd"},
		{"a"},
		{"ac"}}

	expected := []LongestPalindromeOutput{
		{"aaaa"},
		{"bab"},
		{"bb"},
		{"a"},
		{"a"}}

	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		ret := longestPalindrome(input.str)
		assert.Equal(t, expected[i].ret, ret)
	}

}
