package Hot100

// 200. 岛屿数量
// 给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
//
// 岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
//
// 此外，你可以假设该网格的四条边均被水包围。
//
//
//
// 示例 1：
//
// 输入：grid = [
//   ["1","1","1","1","0"],
//   ["1","1","0","1","0"],
//   ["1","1","0","0","0"],
//   ["0","0","0","0","0"]
// ]
// 输出：1
// 示例 2：
//
// 输入：grid = [
//   ["1","1","0","0","0"],
//   ["1","1","0","0","0"],
//   ["0","0","1","0","0"],
//   ["0","0","0","1","1"]
// ]
// 输出：3
//
//
// 提示：
//
// m == grid.length
// n == grid[i].length
// 1 <= m, n <= 300
// grid[i][j] 的值为 '0' 或 '1'

// todo DFS，遍历过的位置由'1'置为'0'

// todo BFS，遍历过的位置由'1'置为'0'

//todo 并查集
func numIslands(grid [][]byte) int {
	ret := 0
	m, n := len(grid), len(grid[0])
	dp := make([][]bool, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]bool, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if grid[i-1][j-1] == '0' {
				continue
			}
			if !dp[i-1][j] && !dp[i][j-1] {
				ret++
			}
			dp[i][j] = true
		}
	}
	return ret
}

// 动态规划（“工”形结构未破解）
func numIslands1(grid [][]byte) int {
	ret := 0
	m, n := len(grid), len(grid[0])
	dp := make([][]bool, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]bool, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if grid[i-1][j-1] == '0' {
				continue
			}
			if !dp[i-1][j] && !dp[i][j-1] {
				ret++
			}
			dp[i][j] = true
		}
	}
	return ret
}
