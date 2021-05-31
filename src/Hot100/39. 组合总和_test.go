package Hot100

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type CombinationSumInput struct {
	candidates []int
	target     int
}

type CombinationSumOutput struct {
	ret [][]int
}

func TestCombinationSum(t *testing.T) {
	inputs := []CombinationSumInput{
		//{[]int{2, 7, 6, 3, 5, 1}, 9},
		{[]int{2, 3, 6, 7}, 7},
		{[]int{2, 3, 5}, 8},
	}

	expected := []CombinationSumOutput{
		{[][]int{{2, 2, 3}, {7}}},
		{[][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}}},
	}

	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		ret := combinationSum(input.candidates, input.target)
		assert.Equal(t, expected[i].ret, ret)
	}

}
