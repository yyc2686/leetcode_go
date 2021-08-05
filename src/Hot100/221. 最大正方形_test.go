package Hot100

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type MaximalSquareInput struct {
	matrix [][]byte
}

type MaximalSquareOutput struct {
	ret int
}

func TestMaximalSquare(t *testing.T) {
	inputs := []MaximalSquareInput{
		{[][]byte{{'0'}, {'1'}}},
		{[][]byte{{'1', '0', '1', '0', '0'}, {'1', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'}}},
		{[][]byte{{'0', '1'}, {'1', '0'}}},
		{[][]byte{{'0'}}},
	}

	expected := []MaximalSquareOutput{
		{1},
		{4},
		{1},
		{0},
	}

	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		ret := maximalSquare(input.matrix)
		assert.Equal(t, expected[i].ret, ret)
	}

}
