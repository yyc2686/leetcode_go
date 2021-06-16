package Hot100

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type MinWindowInput struct {
	s string
	t string
}

type MinWindowOutput struct {
	ret string
}

func TestMinWindow(t *testing.T) {
	inputs := []MinWindowInput{
		{"a", "b"},
		{"a", "aa"},
		{"ADOBECODEBANC", "ABC"},
		{"a", "a"},
	}

	expected := []MinWindowOutput{
		{""},
		{""},
		{"BANC"},
		{"a"},
	}

	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		ret := minWindow(input.s, input.t)
		assert.Equal(t, expected[i].ret, ret)
	}

}
