package Hot100

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TwoSumInput struct {
	nums   []int
	target int
}

type TwoSumOutput struct {
	ret []int
}

func TestTwoSum(t *testing.T) {
	inputs := []TwoSumInput{
		{[]int{2, 7, 11, 15}, 9},
		{[]int{3, 2, 4}, 6},
		{[]int{3, 3}, 6}}
	expected := []TwoSumOutput{
		{[]int{0, 1}},
		{[]int{1, 2}},
		{[]int{0, 1}}}

	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		ret := twoSum(input.nums, input.target)
		assert.Equal(t, expected[i].ret, ret)
	}

}
