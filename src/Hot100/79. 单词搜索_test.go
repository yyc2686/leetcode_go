package Hot100

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type ExistInput struct {
	board [][]byte
	word  string
}

type ExistOutput struct {
	ret bool
}

func TestExist(t *testing.T) {
	inputs := []ExistInput{
		{[][]byte{{'A', 'B'}, {'C', 'D'}}, "DBAC"},
		{[][]byte{{'A', 'A'}}, "AAA"},
		{[][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCCED"},
		{[][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "SEE"},
		{[][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCB"},
	}

	expected := []ExistOutput{
		{true},
		{false},
		{true},
		{true},
		{false},
	}

	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		ret := exist(input.board, input.word)
		assert.Equal(t, expected[i].ret, ret)
	}

}
