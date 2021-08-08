package Hot100

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type LowestCommonAncestorInput struct {
	root, p, q *TreeNode
}

type LowestCommonAncestorOutput struct {
	ret *TreeNode
}

func TestLowestCommonAncestor(t *testing.T) {
	root1 := ConvertSliceToTreeNode([]int{3, 5, 1, 6, 2, 0, 8, -1, -1, 7, 4})
	root2 := ConvertSliceToTreeNode([]int{1, 2})
	inputs := []LowestCommonAncestorInput{
		{root1, root1.Left, root1.Right},
		{root1, root1.Left, root1.Left.Right.Right},
		{root2, root2, root2.Left},
	}

	expected := []LowestCommonAncestorOutput{
		{root1},
		{root1.Left},
		{root2},
	}

	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		ret := lowestCommonAncestor(input.root, input.p, input.q)
		assert.Equal(t, expected[i].ret, ret)
	}

}
