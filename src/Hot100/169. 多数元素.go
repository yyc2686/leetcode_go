package Hot100

//169. 多数元素
//给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。
//
//你可以假设数组是非空的，并且给定的数组总是存在多数元素。
//
//
//
//示例 1：
//
//输入：[3,2,3]
//输出：3
//示例 2：
//
//输入：[2,2,1,1,1,2,2]
//输出：2
//
//
//进阶：
//
//尝试设计时间复杂度为 O(n)、空间复杂度为 O(1) 的算法解决此问题。

// 哈希表法
func majorityElement(nums []int) int {
	m := make(map[int]int, 0)
	size := len(nums)
	for _, val := range nums {
		m[val]++
		if m[val] > size/2 {
			return val
		}
	}
	return -1
}

// 摩尔投票法（更新计数）（空间优化）
func majorityElement2(nums []int) int {
	count := 1
	size := len(nums)
	for i := 1; i < size; i++ {
		if nums[i-1] == nums[i] {
			count++
		} else {
			if count == 0 {
				count = 1
			} else {
				nums[i] = nums[i-1]
				count--
			}
		}
	}
	return nums[size-1]
}

// 摩尔投票法（更新计数）
func majorityElement1(nums []int) int {
	count := 1
	ret := nums[0]
	for i := 1; i < len(nums); i++ {
		if ret == nums[i] {
			count++
		} else {
			if count == 0 {
				count = 1
				ret = nums[i]
			} else {
				count--
			}
		}
	}
	return ret
}
