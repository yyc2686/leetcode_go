package Hot100

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type ReverseListInput struct {
	head *ListNode
}

type ReverseListOutput struct {
	ret *ListNode
}

func TestReverseList(t *testing.T) {

	inputs := []ReverseListInput{
		{ConvertSliceToListNode([]int{1, 2, 3, 4, 5})},
		{ConvertSliceToListNode([]int{1, 2})},
		{ConvertSliceToListNode([]int{})},
	}

	expected := []ReverseListOutput{
		{ConvertSliceToListNode([]int{5, 4, 3, 2, 1})},
		{ConvertSliceToListNode([]int{2, 1})},
		{ConvertSliceToListNode([]int{})},
	}

	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		ret := reverseList(input.head)
		assert.Equal(t, expected[i].ret, ret)
	}
}
