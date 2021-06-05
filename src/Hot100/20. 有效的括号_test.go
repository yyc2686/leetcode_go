package Hot100

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type IsValidInput struct {
	s string
}

type IsValidOutput struct {
	ret bool
}

func TestIsValid(t *testing.T) {
	inputs := []IsValidInput{
		{"()"},
		{"()[]{}"},
		{"(]"},
		{"([)]"},
	}

	expected := []IsValidOutput{
		{true},
		{true},
		{false},
		{false},
	}

	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		ret := isValid(input.s)
		assert.Equal(t, expected[i].ret, ret)
	}

}
