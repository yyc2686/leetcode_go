package Hot100

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type NextPermutationInput struct {
	nums []int
}

type NextPermutationOutput struct {
	ret []int
}

func TestNextPermutation(t *testing.T) {
	inputs := []NextPermutationInput{
		{[]int{2, 3, 1}},
		{[]int{1, 3, 2}},
		{[]int{1, 2, 3, 1, 2}},
		{[]int{1, 2, 3}},
		{[]int{3, 2, 1}},
		{[]int{1, 1, 5}},
		{[]int{1}},
	}

	expected := []NextPermutationOutput{
		{[]int{3, 1, 2}},
		{[]int{2, 1, 3}},
		{[]int{1, 2, 3, 2, 1}},
		{[]int{1, 3, 2}},
		{[]int{1, 2, 3}},
		{[]int{1, 5, 1}},
		{[]int{1}},
	}

	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		ret := nextPermutation(input.nums)
		assert.Equal(t, expected[i].ret, ret)
	}

}
