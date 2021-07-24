package Hot100

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type MajorityElementInput struct {
	nums []int
}

type MajorityElementOutput struct {
	ret int
}

func TestMajorityElement(t *testing.T) {

	inputs := []MajorityElementInput{
		{[]int{3, 2, 3}},
		{[]int{2, 2, 1, 1, 1, 2, 2}},
	}

	expected := []MajorityElementOutput{
		{3},
		{2},
	}

	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		ret := majorityElement(input.nums)
		assert.Equal(t, expected[i].ret, ret)
	}
}
