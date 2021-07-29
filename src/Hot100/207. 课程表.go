package Hot100

//207. 课程表
//你这个学期必须选修 numCourses 门课程，记为 0 到 numCourses - 1 。
//
//在选修某些课程之前需要一些先修课程。 先修课程按数组 prerequisites 给出，其中 prerequisites[i] = [ai, bi] ，表示如果要学习课程 ai 则 必须 先学习课程  bi 。
//
//例如，先修课程对 [0, 1] 表示：想要学习课程 0 ，你需要先完成课程 1 。
//请你判断是否可能完成所有课程的学习？如果可以，返回 true ；否则，返回 false 。
//
//
//
//示例 1：
//
//输入：numCourses = 2, prerequisites = [[1,0]]
//输出：true
//解释：总共有 2 门课程。学习课程 1 之前，你需要完成课程 0 。这是可能的。
//示例 2：
//
//输入：numCourses = 2, prerequisites = [[1,0],[0,1]]
//输出：false
//解释：总共有 2 门课程。学习课程 1 之前，你需要先完成课程 0 ；并且学习课程 0 之前，你还应先完成课程 1 。这是不可能的。
//
//
//提示：
//
//1 <= numCourses <= 105
//0 <= prerequisites.length <= 5000
//prerequisites[i].length == 2
//0 <= ai, bi < numCourses
//prerequisites[i] 中的所有课程对 互不相同

// BFS（通过拓扑排序判断此课程安排图是否是有向无环图DAG）
// 删除所有入度为0的点，网络中仍有节点，则有环
// 引入队列/栈，保证所有入度为0的点都能被考察到
func canFinish(numCourses int, prerequisites [][]int) bool {
	// 初始化度分布
	outInfos := make(map[int][]int, 0)
	inDegrees := make(map[int]int, 0)
	for i := 0; i < numCourses; i++ {
		inDegrees[i] = 0
	}
	for _, prerequisite := range prerequisites {
		inDegrees[prerequisite[1]]++
		if v, _ := outInfos[prerequisite[0]]; v == nil {
			outInfos[prerequisite[0]] = []int{prerequisite[1]}
		} else {
			outInfos[prerequisite[0]] = append(outInfos[prerequisite[0]], prerequisite[1])
		}
	}

	queue := []int{}
	for key, val := range inDegrees {
		if val == 0 {
			queue = append(queue, key)
		}
	}

	for len(queue) != 0 {
		key := queue[0]
		queue = queue[1:]
		delete(inDegrees, key)
		for _, node := range outInfos[key] {
			inDegrees[node]--
			if inDegrees[node] == 0 {
				queue = append(queue, node)
			}
		}
	}

	return len(inDegrees) == 0
}

// todo DFS
func canFinish1(numCourses int, prerequisites [][]int) bool {
	return true
}
