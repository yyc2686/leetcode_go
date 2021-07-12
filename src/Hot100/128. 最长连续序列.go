package Hot100

import "sort"

//128. 最长连续序列
//给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
//
//
//
//进阶：你可以设计并实现时间复杂度为 O(n) 的解决方案吗？
//
//
//
//示例 1：
//
//输入：nums = [100,4,200,1,3,2]
//输出：4
//解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。
//示例 2：
//
//输入：nums = [0,3,7,2,5,8,4,6,0,1]
//输出：9
//
//
//提示：
//
//0 <= nums.length <= 10^4
//-10^9 <= nums[i] <= 10^9

// 哈希表+动态规划（用哈希表存储每个端点值对应连续区间的长度）
func longestConsecutive(nums []int) int {
	size := len(nums)

	if size == 0 {
		return 0
	}

	m := make(map[int]bool, 0)
	for _, num := range nums {
		m[num] = true // 无需判断，直接赋值
	}

	ret := 1
	dp := make(map[int]int, 0)
	for key := range m { // 遍历key时可省略value
		left, right := 0, 0
		if _, ok := dp[key-1]; ok {
			left = dp[key-1]
		}
		if _, ok := dp[key+1]; ok {
			right = dp[key+1]
		}
		cur := left + right + 1
		ret = maxInt(ret, cur)

		dp[key-left], dp[key], dp[key+right] = cur, cur, cur
	}
	return ret
}

// 哈希表（代码优化）
func longestConsecutive3(nums []int) int {
	/*算法思路：
		求一组数中的最长连续序列长度，可以遍历每个数，查询其最左与最右数
		使用哈希表保存数信息（去重）
	分析：最坏情况：O(n^2)
	优化：只有当数为最左元素时才寻找其最右数
	*/

	size := len(nums)

	if size == 0 {
		return 0
	}

	m := make(map[int]bool, 0)
	for _, num := range nums {
		m[num] = true // 无需判断，直接赋值
	}

	ret := 1
	for key := range m { // 遍历key时可省略value
		if !m[key-1] { // 无需赋值，直接判断
			cur := 1
			key++
			for {
				if m[key] { // 无需赋值，直接判断
					cur++
					key++
				} else {
					ret = maxInt(ret, cur)
					break
				}
			}

		}
	}
	return ret
}

// 哈希表
func longestConsecutive2(nums []int) int {
	/*算法思路：
		求一组数中的最长连续序列长度，可以遍历每个数，查询其最左与最右数
		使用哈希表保存数信息（去重）
	分析：最坏情况：O(n^2)
	优化：只有当数为最左元素时才寻找其最右数
	*/

	size := len(nums)

	if size == 0 {
		return 0
	}

	m := make(map[int]bool, 0)
	for _, num := range nums {
		m[num] = true
	}

	ret := 1
	for key := range m {
		if !m[key-1] {
			cur := 1
			key++
			for {
				if m[key] {
					cur++
					key++
				} else {
					ret = maxInt(ret, cur)
					break
				}
			}

		}
	}
	return ret
}

// 排序+模拟（时间复杂度为 O(nlogn) ）
func longestConsecutive1(nums []int) int {
	size := len(nums)

	if size == 0 {
		return 0
	}

	// 升序排列
	sort.Slice(nums, func(i, j int) bool { return nums[i] <= nums[j] })

	// 从左往右遍历更新最长连续序列
	ret := 1
	cur := 1
	for i := 1; i < size; i++ {
		if nums[i] == nums[i-1] {
			continue
		} else if nums[i] == nums[i-1]+1 {
			cur++
		} else {
			ret = maxInt(ret, cur)
			cur = 1
		}
	}
	return maxInt(ret, cur)
}
