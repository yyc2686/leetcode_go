package Hot100

/*对比
11. 盛最多水的容器    64.1%    中等【双指针，单调栈不可行】
42. 接雨水	55.7%	困难【双指针+递归，动态规划，双指针，栈】
84. 柱状图中最大的矩形	42.9%	困难【单调栈（单减栈）】
*/

//84. 柱状图中最大的矩形
//给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
//
//求在该柱状图中，能够勾勒出来的矩形的最大面积。
//
//
//
//
//
//以上是柱状图的示例，其中每个柱子的宽度为 1，给定的高度为 [2,1,5,6,2,3]。
//
//
//
//
//
//图中阴影部分为所能勾勒出的最大矩形面积，其面积为 10 个单位。
//
//
//
//示例:
//
//输入: [2,1,5,6,2,3]
//输出: 10

// 【空间优化（逻辑栈）】单调栈（单减栈，元素小于栈顶元素时出栈）（下标入栈）（触发出栈，元素入栈时下标为最后出栈元素下标）
func largestRectangleArea(heights []int) int {
	ret := 0

	heights = append(heights, -1)

	stack := []int{0}

	for i := 1; i < len(heights); i++ {
		top := -1
		for len(stack) != 0 && heights[i] < heights[stack[len(stack)-1]] {
			top = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ret = maxInt((i-top)*heights[top], ret)
		}
		if top != -1 {
			heights[top] = heights[i]
			stack = append(stack, top)
		} else {
			stack = append(stack, i)
		}
	}

	return ret
}

// 单调栈（单减栈，元素小于栈顶元素时出栈）（下标入栈）（触发出栈，元素入栈时下标为最后出栈元素下标）
func largestRectangleArea2(heights []int) int {
	ret := 0

	heights = append(heights, -1)

	stack := NewStack()
	stack.Push(0)

	for i := 1; i < len(heights); i++ {
		top := -1
		for !stack.IsEmpty() && heights[i] < heights[stack.Peek().(int)] {
			top = stack.Pop().(int)
			ret = maxInt((i-top)*heights[top], ret)
		}
		if top != -1 {
			heights[top] = heights[i]
			stack.Push(top)
		} else {
			stack.Push(i)
		}
	}

	return ret
}

// 暴力解法，以每根柱子为长的矩形（91/96 运行超时）
func largestRectangleArea1(heights []int) int {
	ret := 0
	size := len(heights)
	for i, height := range heights {
		left, right := i, i
		for ; left-1 >= 0 && heights[left-1] >= height; left-- {
		}
		for ; right+1 <= size-1 && heights[right+1] >= height; right++ {
		}
		ret = maxInt(ret, height*(right-left+1))
	}
	return ret
}
