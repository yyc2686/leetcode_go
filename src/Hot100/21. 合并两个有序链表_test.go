package Hot100

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type MergeTwoListsInput struct {
	l1 *ListNode
	l2 *ListNode
}

type MergeTwoListsOutput struct {
	ret *ListNode
}

func TestMergeTwoLists(t *testing.T) {
	inputs := []MergeTwoListsInput{
		{ConvertSliceToListNode([]int{1, 2, 4}), ConvertSliceToListNode([]int{1, 3, 4})},
		{ConvertSliceToListNode([]int{}), ConvertSliceToListNode([]int{})},
		{ConvertSliceToListNode([]int{}), ConvertSliceToListNode([]int{0})},
	}

	expected := []MergeTwoListsOutput{
		{ConvertSliceToListNode([]int{1, 1, 2, 3, 4, 4})},
		{ConvertSliceToListNode([]int{})},
		{ConvertSliceToListNode([]int{0})},
	}

	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		ret := mergeTwoLists(input.l1, input.l2)
		assert.Equal(t, expected[i].ret, ret)
	}

}
