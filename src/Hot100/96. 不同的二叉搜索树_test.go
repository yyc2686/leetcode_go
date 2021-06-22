package Hot100

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type NumTreesInput struct {
	n int
}

type NumTreesOutput struct {
	ret int
}

func TestNumTrees(t *testing.T) {
	inputs := []NumTreesInput{
		{3},
		{1},
	}

	expected := []NumTreesOutput{
		{5},
		{1},
	}

	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		ret := numTrees(input.n)
		assert.Equal(t, expected[i].ret, ret)
	}

}
