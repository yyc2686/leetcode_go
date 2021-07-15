package Hot100

import "strings"

//139. 单词拆分
//给定一个非空字符串 s 和一个包含非空单词的列表 wordDict，判定 s 是否可以被空格拆分为一个或多个在字典中出现的单词。
//
//说明：
//
//拆分时可以重复使用字典中的单词。
//你可以假设字典中没有重复的单词。
//示例 1：
//
//输入: s = "leetcode", wordDict = ["leet", "code"]
//输出: true
//解释: 返回 true 因为 "leetcode" 可以被拆分成 "leet code"。
//示例 2：
//
//输入: s = "applepenapple", wordDict = ["apple", "pen"]
//输出: true
//解释: 返回 true 因为 "applepenapple" 可以被拆分成 "apple pen apple"。
//注意你可以重复使用字典中的单词。
//示例 3：
//
//输入: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
//输出: false

// 回溯+备忘录【未完成】
func wordBreak(s string, wordDict []string) bool {
	if s == "" {
		return true
	}
	return false
}

// 动态规划
func wordBreak2(s string, wordDict []string) bool {
	/*
		dp[i]: s[:i]能否拆分成wordDict的组合
		dp[j] = dp[i] && (s[i:j] in wordDict)
	*/
	size := len(s)
	if size == 0 {
		return true
	}

	m := make(map[string]bool, 0)
	for _, word := range wordDict {
		m[word] = true
	}

	dp := make([]bool, size+1)
	dp[0] = true
	for i := 0; i < size; i++ {
		if !dp[i] {
			continue
		}
		for j := i + 1; j <= size; j++ {
			if m[s[i:j]] {
				dp[j] = true
			}
		}
	}
	return dp[size]
}

// 回溯+递归实现+备忘录（空间换时间）【未完成】
//func wordBreak(s string, wordDict []string) bool {
//	size := len(s)
//	dp := make([][]bool, size)
//
//	var helper func(s1 string, wordDict1 []string) bool
//	helper = func(s1 string, wordDict1 []string) bool {
//		if s1 == "" {
//			return true
//		}
//		for _, word := range wordDict1 {
//			index := strings.Index(s, word)
//			if index == -1 {
//				continue
//			}
//
//			if helper(s1[:index], wordDict1) && helper(s1[index+len(word):], wordDict1) {
//				//dp[]
//				return true
//			}
//		}
//		return false
//	}
//
//	helper(s, wordDict)
//	return dp[size-1][size-1]
//}

// 回溯+递归实现（运行超时）（34 / 44 个通过测试用例）
func wordBreak1(s string, wordDict []string) bool {
	if s == "" {
		return true
	}
	for _, word := range wordDict {
		index := strings.Index(s, word)
		if index == -1 {
			continue
		}

		if wordBreak(s[:index], wordDict) && wordBreak(s[index+len(word):], wordDict) {
			return true
		}
	}
	return false
}
