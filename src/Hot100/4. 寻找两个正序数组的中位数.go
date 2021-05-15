package Hot100

//4. 寻找两个正序数组的中位数
//给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。
//
//
//
//示例 1：
//
//输入：nums1 = [1,3], nums2 = [2]
//输出：2.00000
//解释：合并数组 = [1,2,3] ，中位数 2
//示例 2：
//
//输入：nums1 = [1,2], nums2 = [3,4]
//输出：2.50000
//解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5
//示例 3：
//
//输入：nums1 = [0,0], nums2 = [0,0]
//输出：0.00000
//示例 4：
//
//输入：nums1 = [], nums2 = [1]
//输出：1.00000
//示例 5：
//
//输入：nums1 = [2], nums2 = []
//输出：2.00000
//
//
//提示：
//
//nums1.length == m
//nums2.length == n
//0 <= m <= 1000
//0 <= n <= 1000
//1 <= m + n <= 2000
//-106 <= nums1[i], nums2[i] <= 106
//
//
//进阶：你能设计一个时间复杂度为 O(log (m+n)) 的算法解决此问题吗？
//

// 双指针解法（时间：o(m+n)，空间：O(1)）
func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	left := 0
	right := 0
	len1 := len(nums1)
	len2 := len(nums2)
	mid := (len1 + len2) / 2
	is_even := (len1+len2)%2 == 0
	m := map[int]float64{}
	i := 0
	for {
		if left == len1 || right == len2 || i == mid+1 {
			break
		}

		if nums1[left] < nums2[right] {
			if !is_even && i == mid || is_even && (i == mid-1 || i == mid) {
				m[i] = float64(nums1[left])
			}
			left++
		} else {
			if !is_even && i == mid || is_even && (i == mid-1 || i == mid) {
				m[i] = float64(nums2[right])
			}
			right++
		}
		i++
	}

	if left == len1 {
		for ; i <= mid; i++ {
			if !is_even && i == mid || is_even && (i == mid-1 || i == mid) {
				m[i] = float64(nums2[right])
			}
			right++
		}
	} else if right == len2 {
		for ; i <= mid; i++ {
			if !is_even && i == mid || is_even && (i == mid-1 || i == mid) {
				m[i] = float64(nums1[left])
			}
			left++
		}
	}

	if is_even {
		return 0.5 * (m[mid] + m[mid-1])
	} else {
		return m[mid]
	}
}

// 二分法
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1, len2 := len(nums1), len(nums2)
	if len1 > len2 {
		return findMedianSortedArrays(nums2, nums1)
	}

	if len1 == 0 {
		x, y := findMidian(nums2)
		return 0.5 * float64(nums2[x]+nums2[y])
	} else if len2 == 0 {
		x, y := findMidian(nums1)
		return 0.5 * float64(nums1[x]+nums1[y])
	} else if len1 == 1 && len2 == 1 {
		return 0.5 * float64(nums1[0]+nums2[0])
	}

	left1, right1 := findMidian(nums1)
	left2, right2 := findMidian(nums2)

	for {
		idLeft, maxLeft := max(nums1[left1], nums2[left2])
		_, minRight := min(nums1[right1], nums2[right2])

		if maxLeft <= minRight {
			return 0.5 * float64(minRight+maxLeft)
		}

		if idLeft == 1 {
			left1++
			right1++
			left2--
			right2--
		} else {
			left1--
			right1--
			left2++
			right2++
		}

		if left1 == -1 || right1 == len1 {
			if left2 == -1 {
				if right1 == left1 {
					return float64(nums2[right2])
				} else {
					return 0.5 * float64(nums1[left1]+nums2[right2])
				}
			} else if right2 == len2 {
				if right1 == left1 {
					return float64(nums2[left2])
				} else {
					return 0.5 * float64(nums1[right1]+nums2[left2])
				}
			} else {
				if (len1+len2)%2 == 0 {
					right2++
				} else {
					left2++
				}
				return 0.5 * float64(nums2[left2]+nums2[right2])
			}
		}
	}

}
func findMidian(nums []int) (left int, right int) {
	length := len(nums)
	mid := length / 2
	if length%2 == 0 {
		return mid - 1, mid
	} else {
		return mid, mid
	}
}
func max(x int, y int) (int, int) {
	if x > y {
		return 0, x
	} else {
		return 1, y
	}
}
func min(x int, y int) (int, int) {
	if x < y {
		return 0, x
	} else {
		return 1, y
	}
}
