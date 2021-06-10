package Hot100

import "sort"

//56. 合并区间
//以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
//
//
//
//示例 1：
//
//输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
//输出：[[1,6],[8,10],[15,18]]
//解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
//示例 2：
//
//输入：intervals = [[1,4],[4,5]]
//输出：[[1,5]]
//解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。
//
//
//提示：
//
//1 <= intervals.length <= 104
//intervals[i].length == 2
//0 <= starti <= endi <= 104

// 贪心策略
func mergeIntervals(intervals [][]int) [][]int {
	size := len(intervals)
	ret := make([][]int, 0)

	// 将区间按照右端点逆序排序
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][1] > intervals[j][1] })

	// 依次遍历排序后的区间集
	id := 0
	for {
		if id >= size {
			break
		}
		left, right := intervals[id][0], intervals[id][1]
		for {
			id++
			if id >= size || intervals[id][1] < left {
				break
			}
			if left >= intervals[id][0] && left <= intervals[id][1] {
				left = intervals[id][0]
			}
		}
		ret = append([][]int{{left, right}}, ret...)
	}
	return ret
}
