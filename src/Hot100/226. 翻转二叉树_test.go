package Hot100

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type InvertTreeInput struct {
	root *TreeNode
}

type InvertTreeOutput struct {
	ret *TreeNode
}

func TestInvertTree(t *testing.T) {
	inputs := []InvertTreeInput{
		{ConvertSliceToTreeNode([]int{4, 2, 7, 1, 3, 6, 9})},
	}

	expected := []InvertTreeOutput{
		{ConvertSliceToTreeNode([]int{4, 7, 2, 9, 6, 3, 1})},
	}

	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		ret := invertTree(input.root)
		assert.Equal(t, expected[i].ret, ret)
	}

}
