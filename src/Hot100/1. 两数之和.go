package Hot100

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

//1. 两数之和
//给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 的那 两个 整数，并返回它们的数组下标。
//
//你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
//
//你可以按任意顺序返回答案。
//
//
//
//示例 1：
//
//输入：nums = [2,7,11,15], target = 9
//输出：[0,1]
//解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。
//示例 2：
//
//输入：nums = [3,2,4], target = 6
//输出：[1,2]
//示例 3：
//
//输入：nums = [3,3], target = 6
//输出：[0,1]
//
//
//提示：
//
//2 <= nums.length <= 103
//-109 <= nums[i] <= 109
//-109 <= target <= 109
//只会存在一个有效答案

func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for idx, num := range nums {
		if _, ok := m[target-num]; ok {
			return []int{idx, m[target-num]}
		} else {
			m[num] = idx
		}
	}
	return nil
}

type TwoSumInput struct {
	nums   []int
	target int
}

type TwoSumOutput struct {
	ret []int
}

func TestTwoSum(t *testing.T) {
	inputs := []TwoSumInput{{[]int{2, 7, 11, 15}, 9}, {[]int{3, 2, 4}, 6}, {[]int{3, 3}, 6},}
	expected := []TwoSumOutput{{[]int{0, 1}}, {[]int{1, 2}}, {[]int{0, 1}}}

	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		ret := twoSum(input.nums, input.target)
		assert.Equal(t, expected[i], ret)
	}

}
