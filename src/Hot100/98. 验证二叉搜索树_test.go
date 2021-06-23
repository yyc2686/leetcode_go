package Hot100

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type IsValidBSTInput struct {
	root *TreeNode
}

type IsValidBSTOutput struct {
	ret bool
}

func TestIsValidBST(t *testing.T) {
	inputs := []IsValidBSTInput{
		{ConvertSliceToTreeNode([]int{1, -1, 1})},
		{ConvertSliceToTreeNode([]int{5, 4, 6, -1, -1, 3, 7})},
		{ConvertSliceToTreeNode([]int{2, 2, 2})},
		{ConvertSliceToTreeNode([]int{2, 1, 3})},
		{ConvertSliceToTreeNode([]int{5, 1, 4, -1, -1, 3, 6})},
	}

	expected := []IsValidBSTOutput{
		{false},
		{false},
		{false},
		{true},
		{false},
	}

	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		ret := isValidBST(input.root)
		assert.Equal(t, expected[i].ret, ret)
	}

}
