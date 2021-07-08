package Hot100

//124. 二叉树中的最大路径和
//路径 被定义为一条从树中任意节点出发，沿父节点-子节点连接，达到任意节点的序列。同一个节点在一条路径序列中 至多出现一次 。该路径 至少包含一个 节点，且不一定经过根节点。
//
//路径和 是路径中各节点值的总和。
//
//给你一个二叉树的根节点 root ，返回其 最大路径和 。
//
//
//
//示例 1：
//
//
//输入：root = [1,2,3]
//输出：6
//解释：最优路径是 2 -> 1 -> 3 ，路径和为 2 + 1 + 3 = 6
//示例 2：
//
//
//输入：root = [-10,9,20,null,null,15,7]
//输出：42
//解释：最优路径是 15 -> 20 -> 7 ，路径和为 15 + 20 + 7 = 42
//
//
//提示：
//
//树中节点数目范围是 [1, 3 * 104]
//-1000 <= Node.val <= 1000

// 递归实现
func maxPathSum(root *TreeNode) int {
	VAL := INT_MIN
	var helper func(root *TreeNode, val int) int
	helper = func(root *TreeNode, val int) int {
		if root == nil {
			return 0
		}
		left := helper(root.Left, val)
		right := helper(root.Right, val)
		lmr := root.Val + maxInt(0, left) + maxInt(0, right)
		ret := root.Val + maxInt(0, maxInt(left, right))
		VAL = maxInt(VAL, maxInt(lmr, ret))
		return ret
	}

	helper(root, VAL)
	return VAL
}

func maxPathSum1(root *TreeNode) int {
	queue := NewQueue()
	queue.Offer(root)
	ret := root.Val

	for !queue.IsEmpty() {
		size := queue.Size()
		for i := 0; i < size; i++ {
			node := queue.Poll().(*TreeNode)
			ret = maxInt(ret, node.Val)
			var nodeLeftVal int
			if node.Left != nil {
				val := maxInt(node.Left.Val, node.Left.Val+node.Val)
				ret = maxInt(val, ret)
				nodeLeftVal = node.Left.Val
				node.Left.Val = val
				queue.Offer(node.Left)
			}
			if node.Right != nil {
				val := maxInt(maxInt(node.Right.Val, node.Right.Val+node.Val), nodeLeftVal+node.Right.Val+node.Val)
				ret = maxInt(val, ret)
				node.Right.Val = val
				queue.Offer(node.Right)
			}
		}

	}
	return ret
}
