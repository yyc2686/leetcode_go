package Hot100

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type MaxPathSumInput struct {
	root *TreeNode
}

type MaxPathSumOutput struct {
	ret int
}

func TestMaxPathSum(t *testing.T) {
	inputs := []MaxPathSumInput{
		{ConvertSliceToTreeNode([]int{1, 2, 3})},
		{ConvertSliceToTreeNode([]int{-10, 9, 20, -1, -1, 15, 7})},
		{ConvertSliceToTreeNode([]int{5, 4, 8, 11, -1, 13, 4, 7, 2, -1, -1, -1, 1})},
	}

	expected := []MaxPathSumOutput{
		{6},
		{42},
		{48},
	}

	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		ret := maxPathSum(input.root)
		assert.Equal(t, expected[i].ret, ret)
	}
}
