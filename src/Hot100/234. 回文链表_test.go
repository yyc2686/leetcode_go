package Hot100

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type IsPalindromeInput struct {
	head *ListNode
}

type IsPalindromeOutput struct {
	ret bool
}

func TestIsPalindrome(t *testing.T) {

	inputs := []IsPalindromeInput{
		//{ConvertSliceToListNode([]int{1, 1, 1, 2})},
		//{ConvertSliceToListNode([]int{1, 2})},
		{ConvertSliceToListNode([]int{1, 2, 2, 1})},
	}

	expected := []IsPalindromeOutput{
		//{false},
		//{false},
		{true},
	}

	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		ret := isPalindrome(input.head)
		assert.Equal(t, expected[i].ret, ret)
	}
}
