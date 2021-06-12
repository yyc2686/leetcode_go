package Hot100

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type MinPathSumInput struct {
	grid [][]int
}

type MinPathSumOutput struct {
	ret int
}

func TestMinPathSum(t *testing.T) {
	inputs := []MinPathSumInput{
		{[][]int{{1}}},
		{[][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}},
		{[][]int{{1, 2, 3}, {4, 5, 6}}},
	}

	expected := []MinPathSumOutput{
		{1},
		{7},
		{12},
	}

	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		ret := minPathSum(input.grid)
		assert.Equal(t, expected[i].ret, ret)
	}

}
