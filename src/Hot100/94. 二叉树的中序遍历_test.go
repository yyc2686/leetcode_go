package Hot100

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type InorderTraversalInput struct {
	root *TreeNode
}

type InorderTraversalOutput struct {
	ret []int
}

func TestInorderTraversal(t *testing.T) {
	inputs := []InorderTraversalInput{
		{ConvertSliceToTreeNode([]int{2, 3, -1, 1})},
		{ConvertSliceToTreeNode([]int{1, -1, 2, 3})},
		{ConvertSliceToTreeNode([]int{})},
		{ConvertSliceToTreeNode([]int{1})},
		{ConvertSliceToTreeNode([]int{1, 2})},
		{ConvertSliceToTreeNode([]int{1, -1, 2})},
	}

	expected := []InorderTraversalOutput{
		{[]int{1, 3, 2}},
		{[]int{1, 3, 2}},
		{[]int{}},
		{[]int{1}},
		{[]int{2, 1}},
		{[]int{1, 2}},
	}

	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		ret := inorderTraversal(input.root)
		assert.Equal(t, expected[i].ret, ret)
	}

}
