package Hot100

//79. 单词搜索
//给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；否则，返回 false 。
//
//单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。
//
//
//
//示例 1：
//
//
//输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
//输出：true
//示例 2：
//
//
//输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "SEE"
//输出：true
//示例 3：
//
//
//输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCB"
//输出：false
//
//
//提示：
//
//m == board.length
//n = board[i].length
//1 <= m, n <= 6
//1 <= word.length <= 15
//board 和 word 仅由大小写英文字母组成
//
//
//进阶：你可以使用搜索剪枝的技术来优化解决方案，使其在 board 更大的情况下可以更快解决问题？

// DFS+回溯
func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	size := len(word)

	//if m*n < size {
	//	return false
	//}

	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n) // 记录路径
	}

	var dfs func(x int, y int, index int) bool
	dfs = func(x int, y int, index int) bool {
		if index == size {
			return true
		}

		if x-1 >= 0 && board[x-1][y] == word[index] && !visited[x-1][y] { // 向上走
			visited[x-1][y] = true
			if dfs(x-1, y, index+1) {
				return true
			}
			visited[x-1][y] = false
		}
		if x+1 < m && board[x+1][y] == word[index] && !visited[x+1][y] { // 向下走
			visited[x+1][y] = true
			if dfs(x+1, y, index+1) {
				return true
			}
			visited[x+1][y] = false
		}
		if y-1 >= 0 && board[x][y-1] == word[index] && !visited[x][y-1] { // 向左走
			visited[x][y-1] = true
			if dfs(x, y-1, index+1) {
				return true
			}
			visited[x][y-1] = false
		}
		if y+1 < n && board[x][y+1] == word[index] && !visited[x][y+1] { // 向右走
			visited[x][y+1] = true
			if dfs(x, y+1, index+1) {
				return true
			}
			visited[x][y+1] = false
		}
		return false
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == word[0] {
				visited[i][j] = true
				if dfs(i, j, 1) {
					return true
				}
				visited[i][j] = false
			}
		}
	}
	return false
}
