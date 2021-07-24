package Hot100

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type GetIntersectionNodeInput struct {
	headA, headB *ListNode
}

type GetIntersectionNodeOutput struct {
	ret *ListNode
}

func TestGetIntersectionNode(t *testing.T) {
	expected := []GetIntersectionNodeOutput{
		{ConvertSliceToListNode([]int{8, 4, 5})},
		{ConvertSliceToListNode([]int{2, 4})},
		{ConvertSliceToListNode([]int{})},
	}

	inputs := []GetIntersectionNodeInput{
		{AppendHeadB2TailA(ConvertSliceToListNode([]int{4, 1}), expected[0].ret), AppendHeadB2TailA(ConvertSliceToListNode([]int{5, 0, 1}), expected[0].ret)},
		{AppendHeadB2TailA(ConvertSliceToListNode([]int{0, 9, 1}), expected[1].ret), AppendHeadB2TailA(ConvertSliceToListNode([]int{3}), expected[1].ret)},
		{AppendHeadB2TailA(ConvertSliceToListNode([]int{2, 6, 4}), expected[2].ret), AppendHeadB2TailA(ConvertSliceToListNode([]int{1, 5}), expected[2].ret)},
	}

	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		ret := getIntersectionNode(input.headA, input.headB)
		assert.Equal(t, expected[i].ret, ret)
	}
}
