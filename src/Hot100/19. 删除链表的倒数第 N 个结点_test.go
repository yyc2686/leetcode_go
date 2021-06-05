package Hot100

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type RemoveNthFromEndInput struct {
	head *ListNode
	n    int
}

type RemoveNthFromEndOutput struct {
	ret *ListNode
}

func TestRemoveNthFromEnd(t *testing.T) {
	inputs := []RemoveNthFromEndInput{
		{ConvertSliceToListNode([]int{1, 2, 3, 4, 5}), 2},
		{ConvertSliceToListNode([]int{1}), 1},
		{ConvertSliceToListNode([]int{1, 2}), 1},
	}

	expected := []RemoveNthFromEndOutput{
		{ConvertSliceToListNode([]int{1, 2, 3, 5})},
		{ConvertSliceToListNode([]int{})},
		{ConvertSliceToListNode([]int{1})},
	}

	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		ret := removeNthFromEnd(input.head, input.n)
		assert.Equal(t, expected[i].ret, ret)
	}

}
