package Hot100

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type RroupAnagramsInput struct {
	strs []string
}

type RroupAnagramsOutput struct {
	ret [][]string
}

func TestRroupAnagrams(t *testing.T) {
	inputs := []RroupAnagramsInput{
		{[]string{"eat", "tea", "tan", "ate", "nat", "bat"}},
	}

	expected := []RroupAnagramsOutput{
		{[][]string{{"ate", "eat", "tea"}, {"nat", "tan"}, {"bat"}}},
	}

	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		ret := groupAnagrams(input.strs)
		assert.Equal(t, expected[i].ret, ret)
	}

}
