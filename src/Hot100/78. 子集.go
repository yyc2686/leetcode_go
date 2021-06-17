package Hot100

//78. 子集
//给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
//
//解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。
//
//
//
//示例 1：
//
//输入：nums = [1,2,3]
//输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
//示例 2：
//
//输入：nums = [0]
//输出：[[],[0]]
//
//
//提示：
//
//1 <= nums.length <= 10
//-10 <= nums[i] <= 10
//nums 中的所有元素 互不相同

// 回溯
func subsets(nums []int) [][]int {
	ret := make([][]int, 0)
	size := len(nums)

	var helper func(cur []int, index int)
	helper = func(cur []int, index int) {
		if index == size {
			ret = append(ret, append([]int{}, cur...))
			//ret = append(ret, cur)
			return
		}

		helper(cur, index+1)
		//helper(append(cur, nums[index]), index+1)
		cur = append(cur, nums[index])
		helper(cur, index+1)
		cur = cur[:len(cur)-1]
	}

	helper([]int{}, 0)

	return ret
}
