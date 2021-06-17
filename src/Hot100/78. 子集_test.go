package Hot100

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type SubsetsInput struct {
	nums []int
}

type SubsetsOutput struct {
	ret [][]int
}

func TestSubsets(t *testing.T) {
	inputs := []SubsetsInput{
		{[]int{1, 2, 3}},
		{[]int{0}},
	}

	expected := []SubsetsOutput{
		//{[][]int{{}, {1}, {2}, {3}, {1, 2}, {1, 3}, {2, 3}, {1, 2, 3}}},
		{[][]int{{}, {3}, {2}, {2, 3}, {1}, {1, 3}, {1, 2}, {1, 2, 3}}},
		{[][]int{{}, {0}}},
	}

	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		ret := subsets(input.nums)
		assert.Equal(t, expected[i].ret, ret)
	}

}
