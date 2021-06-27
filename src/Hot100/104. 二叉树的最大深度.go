package Hot100

//104. 二叉树的最大深度
//给定一个二叉树，找出其最大深度。
//
//二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
//
//说明: 叶子节点是指没有子节点的节点。
//
//示例：
//给定二叉树 [3,9,20,null,null,15,7]，
//
//3
/// \
//9  20
///  \
//15   7
//返回它的最大深度 3 。

// BFS+队列/逻辑队列
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ret := 0
	queue := []*TreeNode{root}
	for ; len(queue) > 0; ret++ {
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	return ret
}

// 递归实现(辅助: 求树高)
func maxDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	return maxInt(maxDepth(root.Left), maxDepth(root.Right)) + 1
}
