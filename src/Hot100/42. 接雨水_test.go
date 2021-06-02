package Hot100

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TrapInput struct {
	height []int
}

type TrapOutput struct {
	ret int
}

func TestTrap(t *testing.T) {
	inputs := []TrapInput{
		{[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}},
		{[]int{4, 2, 0, 3, 2, 5}},
	}

	expected := []TrapOutput{
		{6},
		{9},
	}

	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		ret := trap(input.height)
		assert.Equal(t, expected[i].ret, ret)
	}

}
