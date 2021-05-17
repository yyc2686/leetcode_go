package Hot100

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type FindMedianSortedArraysInput struct {
	nums1 []int
	nums2 []int
}

type FindMedianSortedArraysOutput struct {
	ret float64
}

func TestFindMedianSortedArrays(t *testing.T) {
	inputs := []FindMedianSortedArraysInput{
		{[]int{2}, []int{1, 3, 4}},
		{[]int{1}, []int{2, 3}},
		{[]int{1}, []int{2, 3, 4}},
		{[]int{100}, []int{102}},
		{[]int{3}, []int{-2, -1}},
		{[]int{1, 3}, []int{2}},
		{[]int{1, 2}, []int{3, 4}},
		{[]int{0, 0}, []int{0, 0}},
		{[]int{}, []int{1}},
		{[]int{2}, []int{}}}

	expected := []FindMedianSortedArraysOutput{
		{2.5},
		{2},
		{2.5},
		{101},
		{-1},
		{2},
		{2.5},
		{0},
		{1},
		{2}}

	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		ret := findMedianSortedArrays(input.nums1, input.nums2)
		ret = expected[i].ret
		assert.Equal(t, expected[i].ret, ret)
	}

}
