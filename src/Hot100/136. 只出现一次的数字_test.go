package Hot100

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type SingleNumberInput struct {
	nums []int
}

type SingleNumberOutput struct {
	ret int
}

func TestSingleNumber(t *testing.T) {
	inputs := []LongestConsecutiveInput{
		{[]int{2, 2, 1}},
		{[]int{4, 1, 2, 1, 2}},
	}

	expected := []LongestConsecutiveOutput{
		{1},
		{4},
	}

	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		ret := singleNumber(input.nums)
		assert.Equal(t, expected[i].ret, ret)
	}
}
