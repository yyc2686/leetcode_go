package Hot100

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type GenerateParenthesisInput struct {
	n int
}

type GenerateParenthesisOutput struct {
	ret []string
}

func TestGenerateParenthesis(t *testing.T) {
	inputs := []GenerateParenthesisInput{
		{3},
		{1},
	}

	expected := []GenerateParenthesisOutput{
		{[]string{"((()))", "(()())", "(())()", "()(())", "()()()"}},
		{[]string{"()"}},
	}

	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		ret := generateParenthesis(input.n)
		assert.Equal(t, expected[i].ret, ret)
	}

}
