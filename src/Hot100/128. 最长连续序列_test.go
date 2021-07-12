package Hot100

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type LongestConsecutiveInput struct {
	nums []int
}

type LongestConsecutiveOutput struct {
	ret int
}

func TestLongestConsecutive(t *testing.T) {
	inputs := []LongestConsecutiveInput{
		{[]int{100, 4, 200, 1, 3, 2}},
		{[]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}},
	}

	expected := []LongestConsecutiveOutput{
		{4},
		{9},
	}

	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		ret := longestConsecutive(input.nums)
		assert.Equal(t, expected[i].ret, ret)
	}
}
