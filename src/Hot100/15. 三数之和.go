package Hot100

import (
	"bytes"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//15. 三数之和
//给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。
//
//注意：答案中不可以包含重复的三元组。
//
//
//
//示例 1：
//
//输入：nums = [-1,0,1,2,-1,-4]
//输出：[[-1,-1,2],[-1,0,1]]
//示例 2：
//
//输入：nums = []
//输出：[]
//示例 3：
//
//输入：nums = [0]
//输出：[]
//
//
//提示：
//
//0 <= nums.length <= 3000
//-105 <= nums[i] <= 105

// 排序+双指针
// 依次遍历每个元素作为第一元素，双指针分别指向下一个元素与最后一个元素
// 去重：Set去重，(a,b,-(a+b))排序后生成string作为key
// 去重优化：依次遍历每个元素作为第一元素时，重复元素直接跳过；left和right指针与前一个元素相同时也跳过
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	size := len(nums)

	//ret := [][]int{}
	ret := make([][]int, 0)
	for i := 0; i < size-2; i++ {
		if nums[i] > 0 {
			return ret
		}

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, size-1
		target := -nums[i]
		for {
			if left >= right {
				break
			}

			if left > i+1 && nums[left] == nums[left-1] {
				left++
				continue
			}

			if nums[left]+nums[right] > target {
				right--
			} else if nums[left]+nums[right] < target {
				left++
			} else {
				ret = append(ret, []int{nums[i], nums[left], nums[right]})
				left++
				right--
			}
		}

	}

	return ret

}

func threeSum2(nums []int) [][]int {
	sort.Ints(nums)
	size := len(nums)

	ret := [][]int{}
	set := map[string]bool{}
	for i := 0; i < size-2; i++ {
		if nums[i] > 0 {
			return ret
		}

		left, right := i+1, size-1
		for {
			if left >= right {
				break
			}
			if nums[left]+nums[right]+nums[i] == 0 {
				str := strconv.Itoa(nums[i]) + "#" + strconv.Itoa(nums[left]) + "#" + strconv.Itoa(nums[right])
				if _, ok := set[str]; !ok {
					set[str] = true
					ret = append(ret, []int{nums[i], nums[left], nums[right]})
				}
				left++
				right--

			} else if nums[left]+nums[right]+nums[i] > 0 {
				right--
			} else {
				left++
			}

		}

	}

	return ret

}

func countNums(nums []int) (map[int][]int, []int) {
	m := map[int][]int{}
	nums1 := []int{}
	i := 0
	for id, num := range nums {
		if v, ok := m[num]; ok {
			if len(v) < 3 {
				m[num] = append(v, id-i)
				nums1 = append(nums1, num)
			} else {
				i++
			}
		} else {
			m[num] = []int{id - i}
			nums1 = append(nums1, num)
		}
	}
	return m, nums1
}

// 遍历所有(a,b)组合，从map中获取-(a+b)
// map中保存各个元素出现的下标，保证-(a+b)的下标大于b的下标
// 去重：Set去重，(a,b,-(a+b))排序后生成string作为key
func threeSum1(nums []int) [][]int {
	sort.Ints(nums)

	// 统计元素个数
	m, nums := countNums(nums)

	// 遍历所有两数之和，判断相反数
	// 结果去重
	set := map[string]bool{}
	length := len(nums)
	ret := make([][]int, 0)
	for i := 0; i < length-2; i++ {
		for j := i + 1; j < length-1; j++ {
			tmp := nums[i] + nums[j]
			if v, ok := m[-tmp]; ok {
				for _, k := range v {
					if k > j {
						sorted := []int{nums[i], nums[j], -tmp}
						sort.Ints(sorted)
						str := strconv.Itoa(sorted[0]) + "#" + strconv.Itoa(sorted[1]) + "#" + strconv.Itoa(sorted[2])
						if _, ok := set[str]; !ok {
							ret = append(ret, sorted)
							set[str] = true
						}
					}
				}
			}
		}
	}
	return ret

}

func threeSumOperator(nums []int) [][]int {
	sort.Ints(nums)

	// 统计元素个数
	m, nums := countNums(nums)

	// 遍历所有两数之和，判断相反数
	// 结果去重
	set := map[string]bool{}
	length := len(nums)
	ret := make([][]int, 0)
	for i := 0; i < length-2; i++ {
		for j := i + 1; j < length-1; j++ {
			tmp := nums[i] + nums[j]
			if v, ok := m[-tmp]; ok {
				for _, k := range v {
					if k > j {
						sorted := []int{nums[i], nums[j], -tmp}
						sort.Ints(sorted)
						str := strconv.Itoa(sorted[0]) + "#" + strconv.Itoa(sorted[1]) + "#" + strconv.Itoa(sorted[2])
						if _, ok := set[str]; !ok {
							ret = append(ret, sorted)
							set[str] = true
						}
					}
				}
			}
		}
	}
	return ret

}

func threeSumSprintf(nums []int) [][]int {
	sort.Ints(nums)

	// 统计元素个数
	m, nums := countNums(nums)

	// 遍历所有两数之和，判断相反数
	// 结果去重
	set := map[string]bool{}
	length := len(nums)
	ret := make([][]int, 0)
	for i := 0; i < length-2; i++ {
		for j := i + 1; j < length-1; j++ {
			tmp := nums[i] + nums[j]
			if v, ok := m[-tmp]; ok {
				for _, k := range v {
					if k > j {
						sorted := []int{nums[i], nums[j], -tmp}
						sort.Ints(sorted)
						str := fmt.Sprintf("%v#%v#%v", sorted[0], sorted[1], sorted[2])
						if _, ok := set[str]; !ok {
							ret = append(ret, sorted)
							set[str] = true
						}
					}
				}
			}
		}
	}
	return ret

}

func threeSumJoin(nums []int) [][]int {
	sort.Ints(nums)

	// 统计元素个数
	m, nums := countNums(nums)

	// 遍历所有两数之和，判断相反数
	// 结果去重
	set := map[string]bool{}
	length := len(nums)
	ret := make([][]int, 0)
	for i := 0; i < length-2; i++ {
		for j := i + 1; j < length-1; j++ {
			tmp := nums[i] + nums[j]
			if v, ok := m[-tmp]; ok {
				for _, k := range v {
					if k > j {
						sorted := []int{nums[i], nums[j], -tmp}
						sort.Ints(sorted)
						str := strings.Join([]string{strconv.Itoa(sorted[0]), strconv.Itoa(sorted[1]), strconv.Itoa(sorted[2])}, "#")
						//str := fmt.Sprintf("%v#%v#%v", sorted[0], sorted[1], sorted[2])
						if _, ok := set[str]; !ok {
							ret = append(ret, sorted)
							set[str] = true
						}
					}
				}
			}
		}
	}
	return ret

}

func threeSumBuffer(nums []int) [][]int {
	sort.Ints(nums)

	// 统计元素个数
	m, nums := countNums(nums)

	// 遍历所有两数之和，判断相反数
	// 结果去重
	set := map[string]bool{}
	length := len(nums)
	ret := make([][]int, 0)
	for i := 0; i < length-2; i++ {
		for j := i + 1; j < length-1; j++ {
			tmp := nums[i] + nums[j]
			if v, ok := m[-tmp]; ok {
				for _, k := range v {
					if k > j {
						sorted := []int{nums[i], nums[j], -tmp}
						sort.Ints(sorted)
						str := BufferWriteString([]string{strconv.Itoa(sorted[0]), strconv.Itoa(sorted[1]), strconv.Itoa(sorted[2])}, "#")
						if _, ok := set[str]; !ok {
							ret = append(ret, sorted)
							set[str] = true
						}
					}
				}
			}
		}
	}
	return ret

}

func BufferWriteString(strs []string, seq string) string {
	var buffer bytes.Buffer
	for i := 0; i < len(strs)-1; i++ {
		buffer.WriteString(strs[i])
		buffer.WriteString(seq)
	}
	buffer.WriteString(strs[len(strs)-1])
	return buffer.String()
}
