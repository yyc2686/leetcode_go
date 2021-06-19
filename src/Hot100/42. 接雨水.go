package Hot100

/*对比
11. 盛最多水的容器    64.1%    中等【双指针，单调栈不可行】
42. 接雨水	55.7%	困难【双指针+递归，动态规划，双指针，栈】
84. 柱状图中最大的矩形	42.9%	困难【单调栈（单减栈）】
*/

//42. 接雨水
//给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
//
//
//
//示例 1：
//
//
//
//输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
//输出：6
//解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。
//示例 2：
//
//输入：height = [4,2,0,3,2,5]
//输出：9
//
//
//提示：
//
//n == height.length
//0 <= n <= 3 * 104
//0 <= height[i] <= 105

// 动态规划
func trap(height []int) int {
	size := len(height)
	if size <= 2 {
		return 0
	}

	//维护两个数组，分别记录每根柱子左边最大和右边最大值
	leftMax, rightMax := make([]int, 0), make([]int, 0)
	leftMax = append(leftMax, height[0])
	rightMax = append(rightMax, height[size-1])

	// 左边最大遍历
	for i := 1; i < size-1; i++ {
		leftMax = append(leftMax, maxInt(height[i], leftMax[i-1]))
	}

	// 右边最大遍历
	for i := 1; i < size-1; i++ {
		rightMax = append(rightMax, maxInt(height[size-i-1], rightMax[i-1]))
	}

	// 统计每根柱子上的蓄水量
	ret := 0
	for i := 1; i < size-1; i++ {
		ret += minInt(leftMax[i], rightMax[size-i-1]) - height[i]
	}
	return ret
}

// 递归+双指针(优化空间)
func trap2(height []int) int {

	var helper func(l int, r int) int
	helper = func(l int, r int) int {
		if r-l <= 1 {
			return 0
		}

		firstHeightId, secondHeightId := l, l+1
		if height[firstHeightId] < height[secondHeightId] {
			firstHeightId, secondHeightId = secondHeightId, firstHeightId
		}

		for i := l + 2; i <= r; i++ {
			if height[i] > height[firstHeightId] {
				firstHeightId, secondHeightId = i, firstHeightId
			} else if height[i] > height[secondHeightId] {
				secondHeightId = i
			}
		}

		var left int
		var right int
		if firstHeightId < secondHeightId {
			left, right = firstHeightId, secondHeightId
		} else {
			left, right = secondHeightId, firstHeightId
		}

		ret := 0
		for i := left + 1; i < right; i++ {
			ret += height[secondHeightId] - height[i]
		}

		return ret + helper(l, left) + helper(right, r)
	}
	return helper(0, len(height)-1)

}

// 递归+双指针
func trap1(height []int) int {
	size := len(height)
	if size <= 2 {
		return 0
	}

	firstHeightId, secondHeightId := 0, 1
	if height[firstHeightId] < height[secondHeightId] {
		firstHeightId, secondHeightId = secondHeightId, firstHeightId
	}

	for i := 2; i < size; i++ {
		if height[i] > height[firstHeightId] {
			firstHeightId, secondHeightId = i, firstHeightId
		} else if height[i] > height[secondHeightId] {
			secondHeightId = i
		}
	}

	var left int
	var right int
	if firstHeightId < secondHeightId {
		left, right = firstHeightId, secondHeightId
	} else {
		left, right = secondHeightId, firstHeightId
	}

	ret := 0
	for i := left + 1; i < right; i++ {
		ret += height[secondHeightId] - height[i]
	}

	return ret + trap1(height[:left+1]) + trap1(height[right:])
}
