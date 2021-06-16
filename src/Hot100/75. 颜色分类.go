package Hot100

//75. 颜色分类
//给定一个包含红色、白色和蓝色，一共 n 个元素的数组，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。
//
//此题中，我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。
//
//
//
//示例 1：
//
//输入：nums = [2,0,2,1,1,0]
//输出：[0,0,1,1,2,2]
//示例 2：
//
//输入：nums = [2,0,1]
//输出：[0,1,2]
//示例 3：
//
//输入：nums = [0]
//输出：[0]
//示例 4：
//
//输入：nums = [1]
//输出：[1]
//
//
//提示：
//
//n == nums.length
//1 <= n <= 300
//nums[i] 为 0、1 或 2
//
//
//进阶：
//
//你可以不使用代码库中的排序函数来解决这道题吗？
//你能想出一个仅使用常数空间的一趟扫描算法吗？

// 双指针（同时将0和1调整至slice开头）
func sortColors(nums []int) []int {
	size := len(nums)
	p0, p1 := 0, 0

	for i := 0; i < size; i++ {
		if nums[i] == 0 {
			nums[i], nums[p0] = nums[p0], nums[i]
			if p0 < p1 {
				nums[i], nums[p1] = nums[p1], nums[i]
			}
			p0++
			p1++
		} else if nums[i] == 1 {
			nums[i], nums[p1] = nums[p1], nums[i]
			p1++
		}
	}

	return nums
}

// 单指针（分别将0和1调整至slice开头）（两次遍历（非一趟遍历））
func sortColors3(nums []int) []int {
	var helper func(colors []int, target int) (countTarget int)
	helper = func(colors []int, target int) (countTarget int) {
		for i, color := range colors {
			if color == target {
				colors[countTarget], colors[i] = colors[i], colors[countTarget]
				countTarget++
			}
		}
		return countTarget
	}

	countTarget := helper(nums, 0)
	helper(nums[countTarget:], 1)
	return nums
}

// 计数法（记录nums中0，1，2的个数）
func sortColors2(nums []int) []int {
	size := len(nums)
	zero, one := 0, 0
	for _, num := range nums {
		if num == 0 {
			zero++
		} else if num == 1 {
			one++
		}
	}

	i := 0
	for ; i < zero; i++ {
		nums[i] = 0
	}
	for ; i < zero+one; i++ {
		nums[i] = 1
	}
	for ; i < size; i++ {
		nums[i] = 2
	}

	return nums
}

// 冒泡排序（常数空间+二重循环（非一趟遍历））
func sortColors1(nums []int) []int {
	size := len(nums)
	for i := 0; i < size-1; i++ {
		flag := false
		for j := 0; j < size-1; j++ {
			if nums[j+1] < nums[j] {
				nums[j+1], nums[j] = nums[j], nums[j+1]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
	return nums
}
