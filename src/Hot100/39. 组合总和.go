package Hot100

//39. 组合总和
//给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
//
//candidates 中的数字可以无限制重复被选取。
//
//说明：
//
//所有数字（包括 target）都是正整数。
//解集不能包含重复的组合。
//示例 1：
//
//输入：candidates = [2,3,6,7], target = 7,
//所求解集为：
//[
//[7],
//[2,2,3]
//]
//示例 2：
//
//输入：candidates = [2,3,5], target = 8,
//所求解集为：
//[
//[2,2,2,2],
//[2,3,3],
//[3,5]
//]
//
//
//提示：
//
//1 <= candidates.length <= 30
//1 <= candidates[i] <= 200
//candidate 中的每个元素都是独一无二的。
//1 <= target <= 500

// 回溯法，内置函数
func combinationSum2(candidates []int, target int) [][]int {
	ret := make([][]int, 0)
	cur := make([]int, 0)

	var dfs func(target int, id int)
	dfs = func(target int, id int) {
		if target == 0 {
			ret = append(ret, append([]int{}, cur...))
			return
		}

		for i := id; i < len(candidates); i++ {
			if candidates[i] > target {
				continue
			}
			cur = append(cur, candidates[i])
			dfs(target-candidates[i], i)
			cur = cur[:len(cur)-1]
		}
	}

	dfs(target, 0)
	return ret
}

// 回溯法，外置函数
func combinationSum(candidates []int, target int) [][]int {
	allCombinations = make([][]int, 0)
	backTraceCombinationSum(candidates, []int{}, target, 0)
	return allCombinations
}

var allCombinations [][]int

func backTraceCombinationSum(candidates []int, cur []int, target int, id int) {
	if target == 0 {
		//tmp := []int{}
		//for _, v := range cur {
		//	tmp = append(tmp, v)
		//}
		//allCombinations = append(allCombinations, tmp)
		allCombinations = append(allCombinations, append([]int{}, cur...))
		return
	}

	for i := id; i < len(candidates); i++ {
		if candidates[i] > target {
			continue
		}
		backTraceCombinationSum(candidates, append(cur, candidates[i]), target-candidates[i], i)
	}

}
