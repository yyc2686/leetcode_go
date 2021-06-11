package Hot100

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type MaxSubArrayInput struct {
	nums []int
}

type MaxSubArrayOutput struct {
	ret int
}

func TestMaxSubArray(t *testing.T) {
	inputs := []MaxSubArrayInput{
		{[]int{5, 4, -1, 7, 8}},
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}},
		{[]int{1}},
		{[]int{0}},
		{[]int{-1}},
		{[]int{-100000}},
	}
	expected := []MaxSubArrayOutput{
		{23},
		{6},
		{1},
		{0},
		{-1},
		{-100000},
	}

	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		ret := maxSubArray(input.nums)
		assert.Equal(t, expected[i].ret, ret)
	}

}
