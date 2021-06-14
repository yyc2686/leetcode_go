package Hot100

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type MinDistanceInput struct {
	word1 string
	word2 string
}

type MinDistanceOutput struct {
	ret int
}

func TestMinDistance(t *testing.T) {
	inputs := []MinDistanceInput{
		{"horse", "ros"},
		{"intention", "execution"},
	}

	expected := []MinDistanceOutput{
		{3},
		{5},
	}

	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		ret := minDistance(input.word1, input.word2)
		assert.Equal(t, expected[i].ret, ret)
	}

}
