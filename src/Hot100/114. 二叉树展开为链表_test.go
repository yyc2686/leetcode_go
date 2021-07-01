package Hot100

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type FlattenInput struct {
	root *TreeNode
}

type FlattenOutput struct {
	ret *TreeNode
}

func TestFlatten(t *testing.T) {
	inputs := []FlattenInput{
		{ConvertSliceToTreeNode([]int{1, 2, 5, 3, 4, -1, 6})},
		{ConvertSliceToTreeNode([]int{})},
		{ConvertSliceToTreeNode([]int{0})},
	}

	expected := []FlattenOutput{
		{ConvertSliceToTreeNode([]int{1, -1, 2, -1, 3, -1, 4, -1, 5, -1, 6})},
		{ConvertSliceToTreeNode([]int{})},
		{ConvertSliceToTreeNode([]int{0})},
	}

	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		ret := flatten(input.root)
		assert.Equal(t, expected[i].ret, ret)
	}
}
