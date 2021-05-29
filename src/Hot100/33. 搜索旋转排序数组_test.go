package Hot100

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type SearchInput struct {
	nums   []int
	target int
}

type SearchOutput struct {
	ret int
}

func TestSearch(t *testing.T) {
	inputs := []SearchInput{
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 3},
		{[]int{1}, 0},
	}

	expected := []SearchOutput{
		{4},
		{-1},
		{-1},
	}

	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		ret := search(input.nums, input.target)
		assert.Equal(t, expected[i].ret, ret)
	}

}
