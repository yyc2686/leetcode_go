package Hot100

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type MaxProductInput struct {
	nums []int
}

type MaxProductOutput struct {
	ret int
}

func TestMaxProduct(t *testing.T) {

	inputs := []MaxProductInput{
		{[]int{5, -3, 4, -3}},
		{[]int{2, 3, -2, 4}},
		{[]int{-2, 0, -1}},
	}

	expected := []MaxProductOutput{
		{180},
		{6},
		{0},
	}

	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		ret := maxProduct(input.nums)
		assert.Equal(t, expected[i].ret, ret)
	}
}
