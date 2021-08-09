package Hot100

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type ProductExceptSelfInput struct {
	nums []int
}

type ProductExceptSelfOutput struct {
	ret []int
}

func TestProductExceptSelf(t *testing.T) {

	inputs := []ProductExceptSelfInput{
		{[]int{1, 2, 3, 4}},
	}

	expected := []ProductExceptSelfOutput{
		{[]int{24, 12, 8, 6}},
	}

	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		ret := productExceptSelf(input.nums)
		assert.Equal(t, expected[i].ret, ret)
	}
}
