package Hot100

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type PermuteInput struct {
	nums []int
}

type PermuteOutput struct {
	ret [][]int
}

func TestPermute(t *testing.T) {
	inputs := []PermuteInput{
		{[]int{1, 2, 3}},
		{[]int{0, 1}},
		{[]int{1}},
	}

	expected := []PermuteOutput{
		{[][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}},
		{[][]int{{0, 1}, {1, 0}}},
		{[][]int{{1}}},
	}

	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		ret := permute(input.nums)
		assert.Equal(t, expected[i].ret, ret)
	}

}
