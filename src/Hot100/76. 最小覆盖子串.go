package Hot100

//76. 最小覆盖子串
//给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。
//
//注意：如果 s 中存在这样的子串，我们保证它是唯一的答案。
//
//
//
//示例 1：
//
//输入：s = "ADOBECODEBANC", t = "ABC"
//输出："BANC"
//示例 2：
//
//输入：s = "a", t = "a"
//输出："a"
//
//
//提示：
//
//1 <= s.length, t.length <= 105
//s 和 t 由英文字母组成
//
//
//进阶：你能设计一个在 o(n) 时间内解决此问题的算法吗？

// 双指针/滑动窗口（超出时间限制）（再优化，数组存储）
func minWindow(s string, t string) string {
	size, subSize := len(s), len(t)
	if size < subSize {
		return ""
	}

	// 统计字符个数
	strStatic := func(str string) [58]int {
		arr := [58]int{}
		for i := 0; i < len(str); i++ {
			arr[str[i]-'A']++
		}
		return arr
	}

	// 初始化
	left, right := 0, 0
	l, r := -1, -1
	minSize := size
	subArray := strStatic(t)
	srcArray := [58]int{}

	// 判断当前窗口是否满足条件
	check := func() bool {
		for i := 0; i < 58; i++ {
			if subArray[i] > srcArray[i] {
				return false
			}
		}
		return true
	}

	for {
		// 更新srcArray
		srcArray[s[right]-'A']++

		// 若满足条件，更新窗口和srcMap，缩小左端口
		for ; left <= right && check(); left++ {
			if minSize >= right-left+1 {
				minSize = right - left + 1
				l, r = left, right
			}
			srcArray[s[left]-'A']--
		}

		// 不满足条件时，扩大右端口，更新srcMap
		if right >= size-1 {
			break
		} else {
			right++
		}
	}

	if l != -1 {
		return s[l : r+1]
	} else {
		return ""
	}

}

// 双指针/滑动窗口（超出时间限制）（再优化，代码优化）（哈希表存储）
func minWindow3(s string, t string) string {
	size, subSize := len(s), len(t)
	if size < subSize {
		return ""
	}

	// 统计字符个数
	strStatic := func(str string) map[byte]int {
		m := make(map[byte]int, 0)
		for i := 0; i < len(str); i++ {
			if _, ok := m[str[i]]; ok {
				m[str[i]]++
			} else {
				m[str[i]] = 1
			}
		}
		return m
	}

	// 初始化
	left, right := 0, 0
	l, r := -1, -1
	minSize := size
	subMap := strStatic(t)
	srcMap := make(map[byte]int, 0)

	// 判断当前窗口是否满足条件
	check := func() bool {
		for key, val := range subMap {
			if v, ok := srcMap[key]; !ok || v < val {
				return false
			}
		}
		return true
	}

	for {
		// 更新srcMap
		if _, ok := srcMap[s[right]]; ok {
			srcMap[s[right]]++
		} else {
			srcMap[s[right]] = 1
		}

		// 若满足条件，更新窗口和srcMap，缩小左端口
		for ; left <= right && check(); left++ {
			if minSize >= right-left+1 {
				minSize = right - left + 1
				l, r = left, right
			}
			srcMap[s[left]]--
		}

		// 不满足条件时，扩大右端口，更新srcMap
		if right >= size-1 {
			break
		} else {
			right++
		}
	}

	if l != -1 {
		return s[l : r+1]
	} else {
		return ""
	}

}

// 双指针/滑动窗口（超出时间限制）（优化，无需每次从头统计窗口内字符个数）
func minWindow2(s string, t string) string {
	size, subSize := len(s), len(t)
	if size < subSize {
		return ""
	}

	// 统计字符个数
	var strStatic func(str string) map[byte]int
	strStatic = func(str string) map[byte]int {
		m := make(map[byte]int, 0)
		for i := 0; i < len(str); i++ {
			if _, ok := m[str[i]]; ok {
				m[str[i]]++
			} else {
				m[str[i]] = 1
			}
		}
		return m
	}

	// 初始化
	left, right := 0, 0
	l, r := 0, size-1
	minSize := size
	subMap := strStatic(t)
	srcMap := make(map[byte]int, 0)

	// 判断当前窗口是否满足条件
	var isSubString func(srcMap map[byte]int) bool
	isSubString = func(srcMap map[byte]int) bool {
		for key, val := range subMap {
			if v, ok := srcMap[key]; !ok || v < val {
				return false
			}
		}
		return true
	}

	for {
		if left > right {
			break
		}

		// 更新srcMap
		if _, ok := srcMap[s[right]]; ok {
			srcMap[s[right]]++
		} else {
			srcMap[s[right]] = 1
		}

		// 若满足条件，更新窗口和srcMap，缩小左端口
		for ; left <= right && isSubString(srcMap); left++ {
			if minSize > right-left+1 {
				minSize = right - left + 1
				l, r = left, right
			}
			tmp := s[left]
			srcMap[tmp]--
		}

		// 不满足条件时，扩大右端口，更新srcMap
		if right >= size-1 {
			break
		} else {
			right++
		}
	}
	if isSubString(strStatic(s[l : r+1])) {
		return s[l : r+1]
	} else {
		return ""
	}

}

// 双指针/滑动窗口（超出时间限制）
func minWindow1(s string, t string) string {
	size, subSize := len(s), len(t)
	if size < subSize {
		return ""
	}

	var strStatic func(str string) map[byte]int
	strStatic = func(str string) map[byte]int {
		m := make(map[byte]int, 0)
		for i := 0; i < len(str); i++ {
			if _, ok := m[str[i]]; ok {
				m[str[i]]++
			} else {
				m[str[i]] = 1
			}
		}
		return m
	}

	subMap := strStatic(t)

	var isSubString func(src string) bool
	isSubString = func(src string) bool {
		size1 := len(src)
		if size1 < subSize {
			return false
		}
		srcMap := strStatic(src)
		for key, val := range subMap {
			if v, ok := srcMap[key]; !ok || v < val {
				return false
			}
		}
		return true
	}

	left, right := 0, 0
	l, r := 0, size-1
	minSize := size

	for {
		if left > right || right >= size {
			break
		}

		for ; left <= right && isSubString(s[left:right+1]); left++ {
			if minSize > right-left+1 {
				minSize = right - left + 1
				l, r = left, right
			}
		}

		if right < size {
			right++
		} else {
			break
		}

	}

	if isSubString(s[l : r+1]) {
		return s[l : r+1]
	} else {
		return ""
	}

}
