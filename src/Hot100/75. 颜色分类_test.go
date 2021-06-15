package Hot100

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type SortColorsInput struct {
	nums []int
}

type SortColorsOutput struct {
	ret []int
}

func TestSortColors(t *testing.T) {
	inputs := []SortColorsInput{
		//{[]int{0, 2, 1}},
		{[]int{2, 0, 2, 1, 1, 0}},
		{[]int{2, 0, 1}},
		{[]int{0}},
		{[]int{1}},
	}

	expected := []SortColorsOutput{
		//{[]int{0, 1, 2}},
		{[]int{0, 0, 1, 1, 2, 2}},
		{[]int{0, 1, 2}},
		{[]int{0}},
		{[]int{1}},
	}

	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		ret := sortColors(input.nums)
		assert.Equal(t, expected[i].ret, ret)
	}

}
