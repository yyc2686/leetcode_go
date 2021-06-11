package Hot100

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type UniquePathsInput struct {
	m int
	n int
}

type UniquePathsOutput struct {
	ret int
}

func TestUniquePaths(t *testing.T) {
	inputs := []UniquePathsInput{
		{10, 10},
		{4, 4},
		{3, 7},
		{3, 2},
		{7, 3},
		{3, 3},
	}

	expected := []UniquePathsOutput{
		{48620},
		{20},
		{28},
		{3},
		{28},
		{6},
	}

	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		ret := uniquePaths(input.m, input.n)
		assert.Equal(t, expected[i].ret, ret)
	}

}
