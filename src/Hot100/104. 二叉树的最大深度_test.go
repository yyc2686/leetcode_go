package Hot100

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type MaxDepthInput struct {
	root *TreeNode
}

type MaxDepthOutput struct {
	ret int
}

func TestMaxDepth(t *testing.T) {
	inputs := []MaxDepthInput{
		{ConvertSliceToTreeNode([]int{3, 9, 20, -1, -1, 15, 7})},
	}

	expected := []MaxDepthOutput{
		{3},
	}

	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		ret := maxDepth(input.root)
		assert.Equal(t, expected[i].ret, ret)
	}

}
