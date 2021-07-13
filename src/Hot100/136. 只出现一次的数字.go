package Hot100

//136. 只出现一次的数字
//给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
//
//说明：
//
//你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？
//
//示例 1:
//
//输入: [2,2,1]
//输出: 1
//示例 2:
//
//输入: [4,1,2,1,2]
//输出: 4

// 哈希表计数
func singleNumber(nums []int) int {
	m := make(map[int]int, 0)
	for _, num := range nums {
		m[num]++
	}

	for key, val := range m {
		if val == 1 {
			return key
		}
	}

	return -1
}

// 利用异或特性（空间优化）
func singleNumber2(nums []int) int {
	size := len(nums)
	for i := 1; i < size; i++ {
		nums[i] ^= nums[i-1]
	}
	return nums[size-1]
}

// 利用异或特性
func singleNumber1(nums []int) int {
	ret := 0
	for _, num := range nums {
		ret ^= num
	}
	return ret
}
