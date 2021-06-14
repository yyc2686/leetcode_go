package Hot100

//72. 编辑距离
//给你两个单词 word1 和 word2，请你计算出将 word1 转换成 word2 所使用的最少操作数 。
//
//你可以对一个单词进行如下三种操作：
//
//插入一个字符
//删除一个字符
//替换一个字符
//
//
//示例 1：
//
//输入：word1 = "horse", word2 = "ros"
//输出：3
//解释：
//horse -> rorse (将 'h' 替换为 'r')
//rorse -> rose (删除 'r')
//rose -> ros (删除 'e')
//示例 2：
//
//输入：word1 = "intention", word2 = "execution"
//输出：5
//解释：
//intention -> inention (删除 't')
//inention -> enention (将 'i' 替换为 'e')
//enention -> exention (将 'n' 替换为 'x')
//exention -> exection (将 'n' 替换为 'c')
//exection -> execution (插入 'u')
//
//
//提示：
//
//0 <= word1.length, word2.length <= 500
//word1 和 word2 由小写英文字母组成

// 动态规划（空间优化）（两行）
func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = i
	}

	for i := 1; i <= m; i++ {
		tmp := append([]int{}, dp...)
		dp[0] = i
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[j] = tmp[j-1]
			} else {
				dp[j] = 1 + minInt(minInt(tmp[j], dp[j-1]), tmp[j-1])
			}
		}
	}

	return dp[n]
}

// 动态规划（空间未优化）
func minDistance1(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	dp[0][0] = 0
	for j := 1; j <= n; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= m; i++ {
		dp[i][0] = i
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + minInt(minInt(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1])
			}
		}
	}

	return dp[m][n]
}
