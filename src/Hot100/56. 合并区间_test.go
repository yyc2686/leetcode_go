package Hot100

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type MergeIntervalsInput struct {
	intervals [][]int
}

type MergeIntervalsOutput struct {
	ret [][]int
}

func TestMergeIntervals(t *testing.T) {
	inputs := []MergeIntervalsInput{
		{[][]int{{1, 5}, {2, 4}, {3, 8}, {15, 18}}},
		{[][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}},
		{[][]int{{1, 4}, {4, 5}}},
	}

	expected := []MergeIntervalsOutput{
		{[][]int{{1, 8}, {15, 18}}},
		{[][]int{{1, 6}, {8, 10}, {15, 18}}},
		{[][]int{{1, 5}}},
	}

	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		ret := mergeIntervals(input.intervals)
		assert.Equal(t, expected[i].ret, ret)
	}

}
