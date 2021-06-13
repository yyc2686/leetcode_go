package Hot100

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type ClimbStairsInput struct {
	n int
}

type ClimbStairsOutput struct {
	ret int
}

func TestClimbStairs(t *testing.T) {
	inputs := []ClimbStairsInput{
		{1},
		{2},
		{3},
	}

	expected := []ClimbStairsOutput{
		{1},
		{2},
		{3},
	}

	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		ret := climbStairs(input.n)
		assert.Equal(t, expected[i].ret, ret)
	}

}
