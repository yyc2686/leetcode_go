package Hot100

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type LargestRectangleAreaInput struct {
	heights []int
}

type LargestRectangleAreaOutput struct {
	ret int
}

func TestLargestRectangleArea(t *testing.T) {
	inputs := []LargestRectangleAreaInput{
		{[]int{2, 1, 5, 6, 2, 3}},
	}

	expected := []LargestRectangleAreaOutput{
		{10},
	}

	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		ret := largestRectangleArea(input.heights)
		assert.Equal(t, expected[i].ret, ret)
	}

}
