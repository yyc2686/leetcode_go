package Hot100

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type CanJumpsInput struct {
	nums []int
}

type CanJumpsOutput struct {
	ret bool
}

func TestCanJumps(t *testing.T) {
	inputs := []CanJumpsInput{
		{[]int{0}},
		{[]int{2, 3, 1, 1, 4}},
		{[]int{3, 2, 1, 0, 4}},
	}

	expected := []CanJumpsOutput{
		{true},
		{true},
		{false},
	}

	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		ret := canJump(input.nums)
		assert.Equal(t, expected[i].ret, ret)
	}

}
