package Hot100

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type MaxAreaInput struct {
	height []int
}

type MaxAreaOutput struct {
	ret int
}

func TestMaxArea(t *testing.T) {
	inputs := []MaxAreaInput{
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}},
		{[]int{1, 1}},
		{[]int{4, 3, 2, 1, 4}},
		{[]int{1, 2, 1}}}

	expected := []MaxAreaOutput{
		{49},
		{1},
		{16},
		{2}}

	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		ret := maxArea(input.height)
		assert.Equal(t, expected[i].ret, ret)
	}

}
