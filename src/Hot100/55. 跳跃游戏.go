package Hot100

//55. 跳跃游戏
//给定一个非负整数数组 nums ，你最初位于数组的 第一个下标 。
//
//数组中的每个元素代表你在该位置可以跳跃的最大长度。
//
//判断你是否能够到达最后一个下标。
//
//
//
//示例 1：
//
//输入：nums = [2,3,1,1,4]
//输出：true
//解释：可以先跳 1 步，从下标 0 到达下标 1, 然后再从下标 1 跳 3 步到达最后一个下标。
//示例 2：
//
//输入：nums = [3,2,1,0,4]
//输出：false
//解释：无论怎样，总会到达下标为 3 的位置。但该下标的最大跳跃长度是 0 ， 所以永远不可能到达最后一个下标。
//
//
//提示：
//
//1 <= nums.length <= 3 * 104
//0 <= nums[i] <= 105

// 贪心策略（挨着跳，不断更新最远可达距离）
// 当最远可达距离小于当前位置时，不可达
func canJump(nums []int) bool {
	size := len(nums)
	maxDist := nums[0]
	for i := 1; i < size; i++ {
		if maxDist < i {
			return false
		}
		maxDist = maxInt(maxDist, nums[i]+i)
	}
	return true
}

// 动态规划（从i处出发最远可达/是否可达）
func canJump2(nums []int) bool {
	size := len(nums)

	if size == 1 {
		return true
	}

	dp := make([]bool, 0)
	dp = append(dp, nums[size-1] > 0)
	for i := size - 2; i >= 0; i-- {
		flag := false
		if i+nums[i] >= size-1 {
			dp = append(dp, true)
			continue
		}
		for j := nums[i]; j > 0; j-- {
			if dp[(size-1-i)-j] {
				dp = append(dp, true)
				flag = true
				break
			}
		}
		if flag == false {
			dp = append(dp, false)
		}

	}

	return dp[size-1]
}

// 回溯法（超过时间限制）（深度太大）
func canJump1(nums []int) bool {
	size := len(nums)

	var helper func(id int) bool
	helper = func(id int) bool {
		if id >= size-1 {
			return true
		}

		if nums[id] == 0 {
			return false
		}

		for i := nums[id]; i > 0; i-- {
			if helper(id + i) {
				return true
			}
		}

		return false
	}

	return helper(0)
}
