package Hot100

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type SortListInput struct {
	head *ListNode
}

type SortListOutput struct {
	ret *ListNode
}

func TestSortList(t *testing.T) {

	inputs := []SortListInput{
		{ConvertSliceToListNode([]int{4, 2, 1, 3})},
		{ConvertSliceToListNode([]int{-1, 5, 3, 4, 0})},
		{ConvertSliceToListNode([]int{})},
	}

	expected := []SortListOutput{
		{ConvertSliceToListNode([]int{1, 2, 3, 4})},
		{ConvertSliceToListNode([]int{-1, 0, 3, 4, 5})},
		{ConvertSliceToListNode([]int{})},
	}

	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		ret := sortList(input.head)
		assert.Equal(t, expected[i].ret, ret)
	}
}
