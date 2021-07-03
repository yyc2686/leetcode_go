package Hot100

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type MaxProfitInput struct {
	prices []int
}

type MaxProfitOutput struct {
	ret int
}

func TestMaxProfit(t *testing.T) {
	inputs := []MaxProfitInput{
		{[]int{7, 1, 5, 3, 6, 4}},
		{[]int{7, 6, 4, 3, 1}},
	}

	expected := []MaxProfitOutput{
		{5},
		{0},
	}

	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		ret := maxProfit(input.prices)
		assert.Equal(t, expected[i].ret, ret)
	}
}
