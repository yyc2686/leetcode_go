package Hot100

//221. 最大正方形
//在一个由 '0' 和 '1' 组成的二维矩阵内，找到只包含 '1' 的最大正方形，并返回其面积。
//
//
//
//示例 1：
//
//
//输入：matrix = [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]
//输出：4
//示例 2：
//
//
//输入：matrix = [["0","1"],["1","0"]]
//输出：1
//示例 3：
//
//输入：matrix = [["0"]]
//输出：0
//
//
//提示：
//
//m == matrix.length
//n == matrix[i].length
//1 <= m, n <= 300
//matrix[i][j] 为 '0' 或 '1'

// 动态规划+状态压缩
func maximalSquare(matrix [][]byte) int {
	ret := 0
	m, n := len(matrix), len(matrix[0])

	dp := make([]int, n)
	for j := 0; j < n; j++ {
		if matrix[0][j] == '1' {
			dp[j] = 1
			ret = 1
		}
	}

	for i := 1; i < m; i++ {
		tmp := append([]int{}, dp...)
		for j := 0; j < n; j++ {
			if matrix[i][j] == '0' {
				dp[j] = 0
			} else if j == 0 {
				dp[j] = 1
			} else {
				dp[j] = minInt(minInt(dp[j-1], dp[j]), tmp[j-1]) + 1
			}
			ret = maxInt(ret, dp[j]*dp[j])
		}
	}
	return ret
}

// 动态规划，dp[i][j] = (min(dp[i-1][j-1], dp[i][j-1], dp[i-1][j])+1)^2
func maximalSquare1(matrix [][]byte) int {
	ret := 0
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	for j := 0; j < n; j++ {
		if matrix[0][j] == '1' {
			dp[0][j] = 1
			ret = 1
		}
	}

	for i := 0; i < m; i++ {
		if matrix[i][0] == '1' {
			dp[i][0] = 1
			ret = 1
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == '0' {
				continue
			}
			dp[i][j] = minInt(minInt(dp[i][j-1], dp[i-1][j]), dp[i-1][j-1]) + 1
			ret = maxInt(ret, dp[i][j]*dp[i][j])
		}
	}
	return ret
}
