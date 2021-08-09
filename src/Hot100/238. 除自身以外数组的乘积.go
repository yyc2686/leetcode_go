package Hot100

//238. 除自身以外数组的乘积
//给你一个长度为 n 的整数数组 nums，其中 n > 1，返回输出数组 output ，其中 output[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积。
//
//
//
//示例:
//
//输入: [1,2,3,4]
//输出: [24,12,8,6]
//
//
//提示：题目数据保证数组之中任意元素的全部前缀元素和后缀（甚至是整个数组）的乘积都在 32 位整数范围内。
//
//说明: 请不要使用除法，且在 O(n) 时间复杂度内完成此题。
//
//进阶：
//你可以在常数空间复杂度内完成这个题目吗？（ 出于对空间复杂度分析的目的，输出数组不被视为额外空间。）

// 左右乘积+空间优化
func productExceptSelf(nums []int) []int {
	size := len(nums)
	left := make([]int, size)
	left[0] = 1
	for i := 1; i < size; i++ {
		left[i] = left[i-1] * nums[i-1]
	}

	right := 1
	for i := size - 1; i >= 0; i-- {
		left[i] = left[i] * right
		right *= nums[i]
	}
	return left
}

// 左右乘积（省去分类讨论）
func productExceptSelf3(nums []int) []int {
	size := len(nums)
	left := make([]int, size)
	left[0] = 1
	for i := 1; i < size; i++ {
		left[i] = left[i-1] * nums[i-1]
	}

	right := make([]int, size)
	right[size-1] = 1
	for i := size - 2; i >= 0; i-- {
		right[i] = right[i+1] * nums[i+1]
	}

	for i := 0; i < size; i++ {
		nums[i] = left[i] * right[i]
	}
	return nums
}

// 暴力解法+递归实现，运行超时（19 / 20）
func productExceptSelf2(nums []int) []int {
	product := 1
	count := 0
	for _, t := range nums {
		if t == 0 {
			count++
			continue
		}
		product *= t
	}
	if count > 1 {
		for i := 0; i < len(nums); i++ {
			nums[i] = 0
		}
		return nums
	}
	if count == 1 {
		for i := 0; i < len(nums); i++ {
			if nums[i] == 0 {
				nums[i] = product
			} else {
				nums[i] = 0
			}
		}
		return nums
	}

	var helper func(arrs []int) int
	helper = func(arrs []int) int {
		ret := 1
		for _, arr := range arrs {
			ret *= arr
		}
		return ret
	}
	res := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		arr := append([]int{}, nums...)
		res[i] = helper(append(arr[:i], arr[i+1:]...))
	}
	return res
}

// 使用除法，两趟遍历 + 分类讨论
func productExceptSelf1(nums []int) []int {
	product := 1
	count := 0
	for _, t := range nums {
		if t == 0 {
			count++
			continue
		}
		product *= t
	}
	for i := 0; i < len(nums); i++ {
		if count == 0 {
			nums[i] = product / nums[i]
		} else if count == 1 && nums[i] == 0 {
			nums[i] = product
		} else {
			nums[i] = 0
		}
	}
	return nums
}
