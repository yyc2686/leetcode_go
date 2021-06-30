package Hot100

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type BuildTreeInput struct {
	preorder []int
	inorder  []int
}

type BuildTreeOutput struct {
	ret *TreeNode
}

func TestBuildTree(t *testing.T) {
	inputs := []BuildTreeInput{
		{[]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}},
	}

	expected := []BuildTreeOutput{
		{ConvertSliceToTreeNode([]int{3, 9, 20, -1, -1, 15, 7})},
	}

	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		ret := buildTree(input.preorder, input.inorder)
		assert.Equal(t, expected[i].ret, ret)
	}

}
