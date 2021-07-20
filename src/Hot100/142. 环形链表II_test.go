package Hot100

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type DetectCycleInput struct {
	head *ListNode
}

type DetectCycleOutput struct {
	ret *ListNode
}

func TestDetectCycle(t *testing.T) {

	inputs := []DetectCycleInput{
		{LinkTwoListNode(ConvertSliceToListNode([]int{3, 2, 0, -4}), 3, 1)},
		{LinkTwoListNode(ConvertSliceToListNode([]int{1, 2}), 1, 0)},
		{LinkTwoListNode(ConvertSliceToListNode([]int{1}), -1, 0)},
	}

	expected := []DetectCycleOutput{
		{inputs[0].head.Next},
		{inputs[1].head},
		{nil},
	}

	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		ret := detectCycle(input.head)
		assert.Equal(t, expected[i].ret, ret)
	}
}
