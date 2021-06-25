package Hot100

//102. 二叉树的层序遍历
//给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。
//
//
//
//示例：
//二叉树：[3,9,20,null,null,15,7],
//
//3
/// \
//9  20
///  \
//15   7
//返回其层序遍历结果：
//
//[
//[3],
//[9,20],
//[15,7]
//]

// 使用逻辑队列实现
func levelOrder(root *TreeNode) [][]int {
	ret := make([][]int, 0)
	if root == nil {
		return ret
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		size := len(queue)
		tmp := make([]int, size)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			tmp[i] = node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		ret = append(ret, tmp)
	}
	return ret
}

// 使用队列实现
func levelOrder1(root *TreeNode) [][]int {
	ret := make([][]int, 0)
	if root == nil {
		return ret
	}

	queue := NewQueue()
	queue.Offer(root)

	for !queue.IsEmpty() {
		size := queue.Size()
		tmp := make([]int, size)
		for i := 0; i < size; i++ {
			node := queue.Poll().(*TreeNode)
			tmp[i] = node.Val
			if node.Left != nil {
				queue.Offer(node.Left)
			}
			if node.Right != nil {
				queue.Offer(node.Right)
			}
		}
		ret = append(ret, tmp)
	}
	return ret
}
