package Hot100

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type MaxSlidingWindowInput struct {
	nums []int
	k    int
}

type MaxSlidingWindowOutput struct {
	ret []int
}

func TestMaxSlidingWindow(t *testing.T) {

	inputs := []MaxSlidingWindowInput{
		{[]int{1, 3, -1, -3, 5, 3, 6, 7}, 3},
		{[]int{1}, 1},
		{[]int{1, -1}, 1},
		{[]int{9, 11}, 2},
		{[]int{4, -2}, 2},
	}

	expected := []MaxSlidingWindowOutput{
		{[]int{3, 3, 5, 5, 6, 7}},
		{[]int{1}},
		{[]int{1, -1}},
		{[]int{11}},
		{[]int{4}},
	}

	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		ret := maxSlidingWindow(input.nums, input.k)
		assert.Equal(t, expected[i].ret, ret)
	}
}
