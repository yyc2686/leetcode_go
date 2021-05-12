package Hot100

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type AddTwoNumberInput struct {
	l1 *ListNode
	l2 *ListNode
}

type AddTwoNumberOutput struct {
	ret *ListNode
}

func TestAddTwoNumber(t *testing.T) {
	inputs := []AddTwoNumberInput{
		{ConvertSliceToListNode([]int{2, 4, 9}), ConvertSliceToListNode([]int{5, 6, 4, 9})},
		{ConvertSliceToListNode([]int{2, 4, 3}), ConvertSliceToListNode([]int{5, 6, 4})},
		{ConvertSliceToListNode([]int{0}), ConvertSliceToListNode([]int{0})},
		{ConvertSliceToListNode([]int{9, 9, 9, 9, 9, 9, 9}), ConvertSliceToListNode([]int{9, 9, 9, 9})}}

	expected := []AddTwoNumberOutput{
		{ConvertSliceToListNode([]int{7, 0, 4, 0, 1})},
		{ConvertSliceToListNode([]int{7, 0, 8})},
		{ConvertSliceToListNode([]int{0})},
		{ConvertSliceToListNode([]int{8, 9, 9, 9, 0, 0, 0, 1})}}

	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		ret := addTwoNumbers(input.l1, input.l2)
		assert.Equal(t, expected[i].ret, ret)
	}

}
