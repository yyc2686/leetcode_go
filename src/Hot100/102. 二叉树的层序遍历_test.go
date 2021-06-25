package Hot100

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type LevelOrderInput struct {
	root *TreeNode
}

type LevelOrderOutput struct {
	ret [][]int
}

func TestLevelOrder(t *testing.T) {
	inputs := []LevelOrderInput{
		{ConvertSliceToTreeNode([]int{3, 9, 20, -1, -1, 15, 7})},
	}

	expected := []LevelOrderOutput{
		{[][]int{{3}, {9, 20}, {15, 7}}},
	}

	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		ret := levelOrder(input.root)
		assert.Equal(t, expected[i].ret, ret)
	}

}
