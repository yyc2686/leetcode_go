package Hot100

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type HasCycleInput struct {
	head *ListNode
}

type HasCycleOutput struct {
	ret bool
}

func TestHasCycle(t *testing.T) {

	inputs := []HasCycleInput{
		{LinkTwoListNode(ConvertSliceToListNode([]int{1, 2}), -1, 0)},
		{LinkTwoListNode(ConvertSliceToListNode([]int{}), -1, 0)},
		{LinkTwoListNode(ConvertSliceToListNode([]int{3, 2, 0, -4}), 3, 1)},
		{LinkTwoListNode(ConvertSliceToListNode([]int{1, 2}), 1, 0)},
		{LinkTwoListNode(ConvertSliceToListNode([]int{1}), -1, 0)},
	}

	expected := []HasCycleOutput{
		{false},
		{false},
		{true},
		{true},
		{false},
	}

	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		ret := hasCycle(input.head)
		assert.Equal(t, expected[i].ret, ret)
	}
}
