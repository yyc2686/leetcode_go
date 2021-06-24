package Hot100

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type IsSymmetricInput struct {
	root *TreeNode
}

type IsSymmetricOutput struct {
	ret bool
}

func TestIsSymmetric(t *testing.T) {
	inputs := []IsSymmetricInput{
		{ConvertSliceToTreeNode([]int{1, 2, 2, 3, 4, 4, 3})},
		{ConvertSliceToTreeNode([]int{1, 2, 2, -1, 3, -1, 3})},
	}

	expected := []IsSymmetricOutput{
		{true},
		{false},
	}

	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		ret := isSymmetric(input.root)
		assert.Equal(t, expected[i].ret, ret)
	}

}
