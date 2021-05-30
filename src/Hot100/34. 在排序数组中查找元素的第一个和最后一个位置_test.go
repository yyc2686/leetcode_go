package Hot100

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type SearchRangeInput struct {
	nums   []int
	target int
}

type SearchRangeOutput struct {
	ret []int
}

func TestSearchRange(t *testing.T) {
	inputs := []SearchRangeInput{
		{[]int{2, 2}, 2},
		{[]int{5, 7, 7, 8, 8, 10}, 8},
		{[]int{5, 7, 7, 8, 8, 10}, 6},
		{[]int{}, 0},
	}

	expected := []SearchRangeOutput{
		{[]int{0, 1}},
		{[]int{3, 4}},
		{[]int{-1, -1}},
		{[]int{-1, -1}},
	}

	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		ret := searchRange(input.nums, input.target)
		assert.Equal(t, expected[i].ret, ret)
	}

}
