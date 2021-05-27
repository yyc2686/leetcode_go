package Hot100

import "sort"

//31. 下一个排列
//实现获取 下一个排列 的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。
//
//如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。
//
//必须 原地 修改，只允许使用额外常数空间。
//
//
//
//示例 1：
//
//输入：nums = [1,2,3]
//输出：[1,3,2]
//示例 2：
//
//输入：nums = [3,2,1]
//输出：[1,2,3]
//示例 3：
//
//输入：nums = [1,1,5]
//输出：[1,5,1]
//示例 4：
//
//输入：nums = [1]
//输出：[1]
//
//
//提示：
//
//1 <= nums.length <= 100
//0 <= nums[i] <= 100

// 最后升序位 交换 最后「大数」，升序位后翻转
func nextPermutation(nums []int) []int {
	//先找出最大的索引 k 满足 nums[k] < nums[k+1]，如果不存在，就翻转整个数组；
	//再找出另一个最大索引 l 满足 nums[l] > nums[k]；
	//交换 nums[l] 和 nums[k]；
	//最后翻转 nums[k+1:]。

	size := len(nums)
	k := -1

	for k = size - 2; k >= 0; k-- {
		if nums[k] < nums[k+1] {
			break
		}
	}

	if k >= 0 {
		for l := size - 1; l >= 0; l-- {
			if l != k && nums[l] > nums[k] {
				nums[l], nums[k] = nums[k], nums[l]
				break
			}
		}
	}
	reverseSlice(nums[k+1:])
	return nums
}

func reverseSlice(nums []int) {
	size := len(nums)
	i := 0
	j := size - 1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}

// 首个「小数」交换 最小「大数」+ 最小「大数」后升序
func nextPermutation1(nums []int) []int {
	// 需求：下一个排列
	// 要求1：将后面的「大数」与前面的「小数」交换
	// 从后往前找到首个升序对
	// 从升序位开始向后找到最小的「大数」
	// 要求2：增加的幅度尽可能的小
	// 升序位后的序列升序排列

	size := len(nums)
	var i int

	for i = size - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			break
		}
	}
	if i >= 0 {
		minMaxVal := nums[i+1]
		minMaxId := i + 1
		for j := i + 1; j < size; j++ {
			if nums[j] > nums[i] && nums[j] < minMaxVal {
				minMaxVal = nums[j]
				minMaxId = j
			}
		}

		nums[i], nums[minMaxId] = nums[minMaxId], nums[i]
	}
	sort.Ints(nums[i+1:])
	return nums
}

// 按照最大数拆分，前后段分别处理（失败）
//func nextPermutation(nums []int) []int {
//	// 找到最大元素下标
//	maxId := 0
//	maxVal := -1
//	for id, num := range nums {
//		if num >= maxVal {
//			maxId = id
//			maxVal = num
//		}
//	}
//
//	// 判断后序部分是否降序
//	flag := true
//	for i := maxId; i < len(nums)-1; i++ {
//		if nums[i+1] >= nums[i] {
//			flag = false
//			break
//		}
//	}
//
//	if flag {
//		// 后续部分降序，则前序部分调为下一个组合，后续部分升序
//		sort.Ints(nums[maxId:])
//		if maxId >= 1 {
//			nums[maxId-1], nums[maxId] = nums[maxId], nums[maxId-1]
//		}
//	} else {
//		// 后续部分不降序，则前序部分不变，后续部分调为下一个组合
//		nextPermutation(nums[maxId+1:])
//	}
//	return nums
//}
