package Hot100

//3. 无重复字符的最长子串
//给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
//
//
//
//示例 1:
//
//输入: s = "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
//示例 2:
//
//输入: s = "bbbbb"
//输出: 1
//解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
//示例 3:
//
//输入: s = "pwwkew"
//输出: 3
//解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
//示例 4:
//
//输入: s = ""
//输出: 0
//
//
//提示：
//
//0 <= s.length <= 5 * 104
//s 由英文字母、数字、符号和空格组成
func maxInt(v1 int, v2 int) int {
	if v1 <= v2 {
		return v2
	} else {
		return v1
	}
}
func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}

	ret := 0
	left := 0
	right := 1

	m := map[byte]bool{}
	m[s[left]] = true
	for {
		if v, ok := m[s[right]]; ok && v == true {
			ret = maxInt(ret, right-left)
			for i := left + 1; i <= right; i++ {
				delete(m, s[left])
				if s[left] == s[right] {
					left++
					break
				} else {
					left++
				}
			}
			m[s[left]] = true
		}
		m[s[right]] = true
		right++
		if right == len(s) {
			break
		}
	}

	return maxInt(ret, right-left)
}
