package Hot100

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type LetterCombinationsInput struct {
	digits string
}

type LetterCombinationsOutput struct {
	ret []string
}

func TestLetterCombinations(t *testing.T) {
	inputs := []LetterCombinationsInput{
		{"23"},
		{""},
		{"2"},
	}

	expected := []LetterCombinationsOutput{
		{[]string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
		{[]string{}},
		{[]string{"a", "b", "c"}},
	}

	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		ret := letterCombinations(input.digits)
		assert.Equal(t, expected[i].ret, ret)
	}

}
