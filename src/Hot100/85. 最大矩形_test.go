package Hot100

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type MaximalRectangleInput struct {
	matrix [][]byte
}

type MaximalRectangleOutput struct {
	ret int
}

func TestMaximalRectangle(t *testing.T) {
	inputs := []MaximalRectangleInput{
		{[][]byte{{'1', '0', '1', '0', '0'}, {'1', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'}}},
		{[][]byte{}},
		{[][]byte{{'0'}}},
		{[][]byte{{'1'}}},
		{[][]byte{{'0', '0'}}},
	}

	expected := []MaximalRectangleOutput{
		{6},
		{0},
		{0},
		{1},
		{0},
	}

	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		ret := maximalRectangle(input.matrix)
		assert.Equal(t, expected[i].ret, ret)
	}

}
