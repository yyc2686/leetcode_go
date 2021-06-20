package Hot100

//85. 最大矩形
//给定一个仅包含 0 和 1 、大小为 rows x cols 的二维二进制矩阵，找出只包含 1 的最大矩形，并返回其面积。
//
//
//
//示例 1：
//
//
//输入：matrix = [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]
//输出：6
//解释：最大矩形如上图所示。
//示例 2：
//
//输入：matrix = []
//输出：0
//示例 3：
//
//输入：matrix = [["0"]]
//输出：0
//示例 4：
//
//输入：matrix = [["1"]]
//输出：1
//示例 5：
//
//输入：matrix = [["0","0"]]
//输出：0
//
//
//提示：
//
//rows == matrix.length
//cols == matrix[0].length
//0 <= row, cols <= 200
//matrix[i][j] 为 '0' 或 '1'

// 每一层可以看作是柱状图中的最大矩形
func maximalRectangle(matrix [][]byte) int {
	m := len(matrix)
	if m == 0 {
		return 0
	}
	n := len(matrix[0])
	ret := 0

	heights := make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				heights[j] += 1
			} else {
				heights[j] = 0
			}
		}
		ret = maxInt(ret, largestRectangleArea(heights))
	}

	return ret
}

// 动态规划，dp[i][j]：以(i,j)为右下角的最大矩形的左上角坐标（未实现）
func maximalRectangle1(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])
	ret := 0
	dp := make([][][4]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([][4]int, n)
	}

	// (0,0)初始化
	//if matrix[0][0] == '0' {
	//	dp[0][0] = [4]int{-1, -1}
	//} else {
	//	ret = 1
	//	dp[0][0] = [4]int{0, 0}
	//}

	// 第一行初始化
	for j := 1; j < n; j++ {
		if matrix[0][j] == '0' {
			//dp[0][j] = [4]int{-1, -1}
			continue
		} else if matrix[0][j-1] == '0' {
			ret = maxInt(ret, 1)
			dp[0][j] = [4]int{0, j, 0, j}
		} else {
			ret = maxInt(ret, j-dp[0][j-1][1]+1)
			dp[0][j] = dp[0][j-1]
		}
	}

	// 第一列初始化
	for i := 1; i < m; i++ {
		if matrix[i][0] == '0' {
			//dp[i][0] = [4]int{-1, -1}
			continue
		} else if matrix[i-1][0] == '0' {
			ret = maxInt(ret, 1)
			dp[i][0] = [4]int{i, 0, i, 0}
		} else {
			ret = maxInt(ret, i-dp[i-1][0][0]+1)
			dp[i][0] = dp[i-1][0]
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == '0' {
				//dp[i][j] = [4]int{-1, -1}
				continue
			} else if matrix[i-1][j] == '0' && matrix[i][j-1] == '0' {
				ret = maxInt(ret, 1)
				dp[i][j] = [4]int{i, j, i, j}
			} else if matrix[i-1][j] == '0' && matrix[i][j-1] == '1' {
				ret = maxInt(ret, (i-dp[i][j-1][0]+1)*(j-dp[i][j-1][1]+1))
				dp[i][j] = [4]int{dp[i][j-1][0], dp[i][j-1][1], i, j}
			} else if matrix[i-1][j] == '1' && matrix[i][j-1] == '0' {
				ret = maxInt(ret, (i-dp[i-1][j][2]+1)*(j-dp[i-1][j][3]+1))
				dp[i][j] = [4]int{i, j, dp[i-1][j][2], dp[i-1][j][3]}
			} else {
				tmp1 := (i - dp[i][j-1][0] + 1) * (j - dp[i][j-1][1] + 1)
				tmp2 := (i - dp[i-1][j][2] + 1) * (j - dp[i-1][j][3] + 1)
				if tmp1 > tmp2 {
					dp[i][j] = dp[i][j-1]
					ret = maxInt(ret, tmp1)
				} else {
					dp[i][j] = dp[i-1][j]
					ret = maxInt(ret, tmp2)
				}
			}
		}
	}
	return ret
}
