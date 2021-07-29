package Hot100

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type CanFinishInput struct {
	numCourses    int
	prerequisites [][]int
}

type CanFinishOutput struct {
	ret bool
}

func TestCanFinish(t *testing.T) {

	inputs := []CanFinishInput{
		{5, [][]int{{1, 4}, {2, 4}, {3, 1}, {3, 2}}},
		{2, [][]int{{1, 0}}},
		{2, [][]int{{1, 0}, {0, 1}}},
	}

	expected := []CanFinishOutput{
		{true},
		{true},
		{false},
	}

	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		ret := canFinish(input.numCourses, input.prerequisites)
		assert.Equal(t, expected[i].ret, ret)
	}
}
