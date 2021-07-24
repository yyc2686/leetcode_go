package Hot100

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type RobInput struct {
	nums []int
}

type RobOutput struct {
	ret int
}

func TestRob(t *testing.T) {

	inputs := []RobInput{
		{[]int{1, 2, 3, 1}},
		{[]int{2, 7, 9, 3, 1}},
	}

	expected := []RobOutput{
		{4},
		{12},
	}

	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		ret := rob(input.nums)
		assert.Equal(t, expected[i].ret, ret)
	}
}
