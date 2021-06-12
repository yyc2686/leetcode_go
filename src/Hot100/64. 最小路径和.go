package Hot100

//64. 最小路径和
//给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
//
//说明：每次只能向下或者向右移动一步。
//
//
//
//示例 1：
//
//
//输入：grid = [[1,3,1],[1,5,1],[4,2,1]]
//输出：7
//解释：因为路径 1→3→1→1→1 的总和最小。
//示例 2：
//
//输入：grid = [[1,2,3],[4,5,6]]
//输出：12
//
//
//提示：
//
//m == grid.length
//n == grid[i].length
//1 <= m, n <= 200
//0 <= grid[i][j] <= 100

//bfs: 广度优先搜索（队列）（空间优化）
func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	keyQueue := NewQueue()
	valMap := make(map[[2]int]int, 0)

	keyQueue.Offer([2]int{0, 0})
	valMap[[2]int{0, 0}] = grid[0][0]

	for {
		if keyQueue.IsEmpty() {
			break
		}

		flag := false

		size := keyQueue.Size()
		for i := 0; i < size; i++ {
			key := keyQueue.Poll().([2]int)
			val := valMap[key]

			if key[0] <= m-1 && key[1] < n-1 {
				newKey := [2]int{key[0], key[1] + 1}
				newVal := val + grid[newKey[0]][newKey[1]]
				if oldVal, ok := valMap[newKey]; !ok || newVal < oldVal {
					valMap[newKey] = newVal
					flag = true
				}
			}
			if key[0] < m-1 && key[1] <= n-1 {
				newKey := [2]int{key[0] + 1, key[1]}
				newVal := val + grid[newKey[0]][newKey[1]]
				if oldVal, ok := valMap[newKey]; !ok || newVal < oldVal {
					valMap[newKey] = newVal
					flag = true
				}
			}

			if flag {
				delete(valMap, key)
			} else {
				return valMap[[2]int{m - 1, n - 1}]
			}
		}

		for k, _ := range valMap {
			keyQueue.Offer(k)
		}
	}

	return valMap[[2]int{m - 1, n - 1}]
}

//bfs: 广度优先搜索（队列）（空间未优化）
func minPathSum3(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	keyQueue := NewQueue()
	valMap := make(map[[2]int]int, 0)

	keyQueue.Offer([2]int{0, 0})
	valMap[[2]int{0, 0}] = grid[0][0]

	for {
		if keyQueue.IsEmpty() {
			break
		}

		size := keyQueue.Size()
		for i := 0; i < size; i++ {
			key := keyQueue.Poll().([2]int)
			val := valMap[key]
			//delete(valMap, key)
			if key[0] <= m-1 && key[1] < n-1 {
				newKey := [2]int{key[0], key[1] + 1}
				newVal := val + grid[newKey[0]][newKey[1]]
				if oldVal, ok := valMap[newKey]; !ok || newVal < oldVal {
					keyQueue.Offer(newKey)
					valMap[newKey] = newVal
				}
			}
			if key[0] < m-1 && key[1] <= n-1 {
				newKey := [2]int{key[0] + 1, key[1]}
				newVal := val + grid[newKey[0]][newKey[1]]
				if oldVal, ok := valMap[newKey]; !ok || newVal < oldVal {
					keyQueue.Offer(newKey)
					valMap[newKey] = newVal
				}
			}
		}
	}

	return valMap[[2]int{m - 1, n - 1}]
}

// 动态规划（空间优化）
func minPathSum2(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([]int, n)

	dp[0] = grid[0][0]

	for j := 1; j < n; j++ {
		dp[j] = dp[j-1] + grid[0][j]
	}

	for i := 1; i < m; i++ {
		dp[0] += grid[i][0]
		for j := 1; j < n; j++ {
			dp[j] = minInt(dp[j], dp[j-1]) + grid[i][j]
		}
	}

	return dp[n-1]

}

// 动态规划（空间未优化）
func minPathSum1(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	dp[0][0] = grid[0][0]

	for j := 1; j < n; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}

	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
		for j := 1; j < n; j++ {
			dp[i][j] = minInt(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}

	return dp[m-1][n-1]

}
