package Hot100

//46. 全排列
//给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
//
//
//
//示例 1：
//
//输入：nums = [1,2,3]
//输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
//示例 2：
//
//输入：nums = [0,1]
//输出：[[0,1],[1,0]]
//示例 3：
//
//输入：nums = [1]
//输出：[[1]]
//
//
//提示：
//
//1 <= nums.length <= 6
//-10 <= nums[i] <= 10
//nums 中的所有整数 互不相同

// 回溯法
func permute(nums []int) [][]int {
	ret := [][]int{}

	var helper func(cur []int, selected []int)
	helper = func(cur []int, selected []int) {
		if len(selected) == 0 {
			ret = append(ret, cur)
			return
		}
		for id, num := range selected {
			//tmp := append([]int{}, selected...)
			tmp := make([]int, len(selected))
			copy(tmp, selected)
			helper(append(cur, num), append(tmp[:id], tmp[id+1:]...))
		}
	}

	helper([]int{}, nums)
	return ret
}
