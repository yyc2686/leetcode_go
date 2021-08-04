package Hot100

import "sort"

//215. 数组中的第K个最大元素
//给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。
//
//请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
//
//
//
//示例 1:
//
//输入: [3,2,1,5,6,4] 和 k = 2
//输出: 5
//示例 2:
//
//输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
//输出: 4
//
//
//提示：
//
//1 <= k <= nums.length <= 104
//-104 <= nums[i] <= 104

// todo 基于堆排
func findKthLargest3(nums []int, k int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] >= nums[j]
	})
	return nums[k-1]
}

// 基于快排（应直接使用nums+(i,j)进行数组范围缩小）
func findKthLargest(nums []int, k int) int {
	ret := nums[k-1]
	size := len(nums)
	var quickSort func(x, y int) bool
	quickSort = func(x, y int) bool {
		if x > y {
			return false
		}
		tmp := nums[x]
		cur, i, j := x, x+1, y
		for {
			for j >= x && nums[j] <= tmp {
				j--
			}
			if i > j || j <= x {
				break
			}
			nums[j], nums[cur] = nums[cur], nums[j]
			cur = j
			j--

			for i <= y && nums[i] >= tmp {
				i++
			}
			if i > j || i >= y {
				break
			}
			nums[i], nums[cur] = nums[cur], nums[i]
			cur = i
			i++
		}
		if cur == k-1 {
			ret = nums[cur]
			return true
		} else if cur > k-1 {
			return quickSort(x, cur-1)
		} else {
			return quickSort(cur+1, y)
		}
	}
	quickSort(0, size-1)
	return ret
}

// 排序
func findKthLargest1(nums []int, k int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] >= nums[j]
	})
	return nums[k-1]
}
