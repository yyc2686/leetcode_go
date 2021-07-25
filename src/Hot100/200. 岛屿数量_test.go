package Hot100

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type NumIslandsInput struct {
	ngrid [][]byte
}

type NumIslandsOutput struct {
	ret int
}

func TestNumIslands(t *testing.T) {

	inputs := []NumIslandsInput{
		{[][]byte{{'1', '1', '1'}, {'0', '1', '0'}, {'1', '1', '1'}}},
		{[][]byte{{'1', '1', '1', '1', '0'}, {'1', '1', '0', '1', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '0', '0', '0'}}},
		{[][]byte{{'1', '1', '0', '0', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '1', '0', '0'}, {'0', '0', '0', '1', '1'}}},
	}

	expected := []NumIslandsOutput{
		{1},
		{1},
		{3},
	}

	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		ret := numIslands(input.ngrid)
		assert.Equal(t, expected[i].ret, ret)
	}
}
