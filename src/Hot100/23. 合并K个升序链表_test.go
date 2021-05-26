package Hot100

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type MergeKListsInput struct {
	lists []*ListNode
}

type MergeKListsOutput struct {
	ret *ListNode
}

func TestMergeKLists(t *testing.T) {
	inputs := []MergeKListsInput{
		{[]*ListNode{ConvertSliceToListNode([]int{1}), ConvertSliceToListNode([]int{0})}},
		{[]*ListNode{ConvertSliceToListNode([]int{1, 4, 5}), ConvertSliceToListNode([]int{1, 3, 4}), ConvertSliceToListNode([]int{2, 6})}},
		{[]*ListNode{}},
		{[]*ListNode{ConvertSliceToListNode([]int{})}},
	}

	expected := []MergeKListsOutput{
		{ConvertSliceToListNode([]int{0, 1})},
		{ConvertSliceToListNode([]int{1, 1, 2, 3, 4, 4, 5, 6})},
		{ConvertSliceToListNode([]int{})},
		{ConvertSliceToListNode([]int{})},
	}

	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		ret := mergeKLists(input.lists)
		assert.Equal(t, expected[i].ret, ret)
	}

}
