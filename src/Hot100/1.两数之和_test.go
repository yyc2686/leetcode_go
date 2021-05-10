package Hot100

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//func twoSum(nums []int, target int) []int {
//	m := map[int]int{}
//	for idx, num := range nums {
//		if _, ok := m[target-num]; ok {
//			return []int{m[target-num], idx}
//		} else {
//			m[num] = idx
//		}
//	}
//	return nil
//}

type TwoSumInput struct {
	nums   []int
	target int
}

type TwoSumOutput struct {
	ret []int
}

func TestTwoSum(t *testing.T) {

	//println("Hello world")

	inputs := []TwoSumInput{{[]int{2, 7, 11, 15}, 9}, {[]int{3, 2, 4}, 6}, {[]int{3, 3}, 6}}
	expected := []TwoSumOutput{{[]int{0, 1}}, {[]int{1, 2}}, {[]int{0, 1}}}

	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		ret := twoSum(input.nums, input.target)
		assert.Equal(t, expected[i].ret, ret)
	}

}
