package Hot100

import "math/big"

//62. 不同路径
//一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。
//
//机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。
//
//问总共有多少条不同的路径？
//
//
//
//示例 1：
//
//
//输入：m = 3, n = 7
//输出：28
//示例 2：
//
//输入：m = 3, n = 2
//输出：3
//解释：
//从左上角开始，总共有 3 条路径可以到达右下角。
//1. 向右 -> 向下 -> 向下
//2. 向下 -> 向下 -> 向右
//3. 向下 -> 向右 -> 向下
//示例 3：
//
//输入：m = 7, n = 3
//输出：28
//示例 4：
//
//输入：m = 3, n = 3
//输出：6
//
//
//提示：
//
//1 <= m, n <= 100
//题目数据保证答案小于等于 2 * 109

// 排列组合 C(m+n-2, n-1)（存在误差）
func uniquePaths(m int, n int) int {
	if m < n {
		return uniquePaths(n, m)
	}

	return int(new(big.Int).Binomial(int64(m+n-2), int64(n-1)).Int64())

	//var ret float64 = 1
	//for i := 1; i <= n-1; i++ {
	//	ret *= float64((i + m - 1)) / float64(i)
	//}

	//return int(ret)

}

// 动态规划（优化空间）（使用一行空间+取min(m, n)）（优化动态数组创建）
func uniquePaths4(m int, n int) int {
	if m < n {
		return uniquePaths(n, m)
	}

	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[j] = dp[j] + dp[j-1]
		}
	}

	return dp[n-1]

}

// 动态规划（优化空间）（使用一行空间+取min(m, n)）
func uniquePaths3(m int, n int) int {
	if m < n {
		return uniquePaths(n, m)
	}

	dp := make([]int, 0)
	for i := 0; i < n; i++ {
		dp = append(dp, 1)
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[j] = dp[j] + dp[j-1]
		}
	}

	return dp[n-1]

}

// 动态规划（优化空间）（使用两行空间）
func uniquePaths2(m int, n int) int {
	dp := make([]int, 0)
	for i := 0; i < n; i++ {
		dp = append(dp, 1)
	}

	for i := 1; i < m; i++ {
		tmp := []int{1}
		for j := 1; j < n; j++ {
			tmp = append(tmp, dp[j]+tmp[j-1])
		}
		dp = append([]int{}, tmp...)
	}

	return dp[n-1]

}

// 动态规划（未优化空间）
func uniquePaths1(m int, n int) int {
	dp := make([][]int, 0)

	tmp := make([]int, 0)
	for i := 0; i < n; i++ {
		tmp = append(tmp, 1)
	}
	dp = append(dp, tmp)

	for i := 1; i < m; i++ {
		tmp = []int{1}
		for j := 1; j < n; j++ {
			tmp = append(tmp, dp[i-1][j]+tmp[j-1])
		}
		dp = append(dp, tmp)
	}

	return dp[m-1][n-1]

}
