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

// 并查集（可优化）
type unionFindSet struct {
	ufSet  map[Element]Element
	rank   map[Element]int
	branch int
}

func NewUnionFindSet() *unionFindSet {
	return &unionFindSet{}
}

type UnionFindSet interface {
	Find(e Element) Element     //查找元素对应的首领
	Union(x Element, y Element) // 合并两个元素对应的帮派
}

//查找元素对应的首领
func (entry *unionFindSet) Find(e Element) Element {
	if val, ok := entry.ufSet[e]; !ok {
		entry.ufSet[e] = e
		entry.rank[e]++
		entry.branch++
		return e
	} else if val == e {
		return e
	} else {
		entry.ufSet[e] = entry.Find(entry.ufSet[e])
		return entry.ufSet[e]
	}
}

func (entry *unionFindSet) Union(x Element, y Element) {
	i, j := entry.Find(x), entry.Find(y)
	if i == j {
		return
	}
	if entry.rank[i] < entry.rank[j] {
		entry.ufSet[i] = j
		entry.rank[j]++
	} else {
		entry.ufSet[j] = i
		entry.rank[i]++
	}
	entry.branch--
}

func (entry *unionFindSet) GetBranchNum() int {
	return entry.branch
}

func numIslands(grid [][]byte) int {
	ret := 0
	m, n := len(grid), len(grid[0])

	entry := NewUnionFindSet()
	entry.ufSet = make(map[Element]Element, 0)
	entry.rank = make(map[Element]int, 0)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '0' {
				continue
			}
			x := entry.Find([2]int{i, j})
			var y Element
			// 向右走
			if j+1 < n && grid[i][j+1] == '1' {
				y = entry.Find([2]int{i, j + 1})
				entry.Union(x, y)
			}

			// 向下走
			if i+1 < m && grid[i+1][j] == '1' {
				y = entry.Find([2]int{i + 1, j})
				entry.Union(x, y)
			}
		}
	}

	// 统计连通分支数（优化后）
	return entry.GetBranchNum()

	// 统计连通分支数（优化前）
	ufSetMap := make(map[Element]bool, 0)
	for _, val := range entry.ufSet {
		realVal := entry.Find(val)
		if ufSetMap[realVal] {
			continue
		}
		ret++
		ufSetMap[realVal] = true
	}
	return ret
}

// BFS，遍历过的位置由'1'置为'0'
func numIslands3(grid [][]byte) int {
	ret := 0
	m, n := len(grid), len(grid[0])
	queue := [][2]int{}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '0' {
				continue
			}
			ret++
			queue = append(queue, [2]int{i, j})
			for len(queue) != 0 {
				node := queue[0]
				queue = queue[1:]
				grid[node[0]][node[1]] = '0'
				// 向左走
				if node[1]-1 >= 0 && grid[node[0]][node[1]-1] == '1' {
					queue = append(queue, [2]int{node[0], node[1] - 1})
					grid[node[0]][node[1]-1] = '0'
				}

				// 向右走
				if node[1]+1 < n && grid[node[0]][node[1]+1] == '1' {
					queue = append(queue, [2]int{node[0], node[1] + 1})
					grid[node[0]][node[1]+1] = '0'
				}

				// 向下走
				if node[0]+1 < m && grid[node[0]+1][node[1]] == '1' {
					queue = append(queue, [2]int{node[0] + 1, node[1]})
					grid[node[0]+1][node[1]] = '0'
				}

				// 向上走
				if node[0]-1 >= 0 && grid[node[0]-1][node[1]] == '1' {
					queue = append(queue, [2]int{node[0] - 1, node[1]})
					grid[node[0]-1][node[1]] = '0'
				}

			}
		}
	}

	return ret
}

// DFS，遍历过的位置由'1'置为'0'
func numIslands2(grid [][]byte) int {
	ret := 0
	m, n := len(grid), len(grid[0])

	var helper func(x, y int)
	helper = func(x, y int) {
		if grid[x][y] == '0' {
			return
		}
		grid[x][y] = '0'

		// 向左走
		if y-1 >= 0 {
			helper(x, y-1)
		}

		// 向右走
		if y+1 < n {
			helper(x, y+1)
		}

		// 向下走
		if x+1 < m {
			helper(x+1, y)
		}

		// 向上走
		if x-1 >= 0 {
			helper(x-1, y)
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '0' {
				continue
			}
			helper(i, j)
			ret++
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
