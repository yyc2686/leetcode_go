package Hot100

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type RotateInput struct {
	matrix [][]int
}

type RotateOutput struct {
	ret [][]int
}

func TestRotate(t *testing.T) {
	inputs := []RotateInput{
		{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}},
		{[][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}},
		{[][]int{{1}}},
		{[][]int{{1, 2}, {3, 4}}},
	}

	expected := []RotateOutput{
		{[][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}}},
		{[][]int{{15, 13, 2, 5}, {14, 3, 4, 1}, {12, 6, 8, 9}, {16, 7, 10, 11}}},
		{[][]int{{1}}},
		{[][]int{{3, 1}, {4, 2}}},
	}

	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		ret := rotate(input.matrix)
		assert.Equal(t, expected[i].ret, ret)
	}

}
