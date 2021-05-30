package Hot100

//34. 在排序数组中查找元素的第一个和最后一个位置
//给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
//
//如果数组中不存在目标值 target，返回 [-1, -1]。
//
//进阶：
//
//你可以设计并实现时间复杂度为 O(log n) 的算法解决此问题吗？
//
//
//示例 1：
//
//输入：nums = [5,7,7,8,8,10], target = 8
//输出：[3,4]
//示例 2：
//
//输入：nums = [5,7,7,8,8,10], target = 6
//输出：[-1,-1]
//示例 3：
//
//输入：nums = [], target = 0
//输出：[-1,-1]
//
//
//提示：
//
//0 <= nums.length <= 105
//-109 <= nums[i] <= 109
//nums 是一个非递减数组
//-109 <= target <= 109

// 二分法
func searchRange(nums []int, target int) []int {
	return []int{searchLeftRange(nums, target), searchRightRange(nums, target)}
}

func searchLeftRange(nums []int, target int) int {
	size := len(nums)
	left, right := 0, size-1
	for {
		if left > right {
			break
		}

		mid := left + (right-left)>>1
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			right = mid - 1 // 收缩右侧边界
		}
	}

	if left >= size || nums[left] != target {
		return -1
	}
	return left
}

func searchRightRange(nums []int, target int) int {
	size := len(nums)
	left, right := 0, size-1
	for {
		if left > right {
			break
		}

		mid := left + (right-left)>>1
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1 // 收缩左侧边界
		}
	}

	if right < 0 || nums[right] != target {
		return -1
	}
	return right
}

// 暴力解法
func searchRange1(nums []int, target int) []int {
	size := len(nums)
	left, right := 0, size-1

	for {
		// 边界处理1
		if left > right {
			return []int{-1, -1}
		}
		// 边界处理2
		if nums[left] > target || nums[right] < target {
			return []int{-1, -1}
		}

		if nums[left] == target && nums[right] == target {
			break
		}

		if nums[left] < target {
			left++
		}

		if nums[right] > target {
			right--
		}
	}

	return []int{left, right}
}
