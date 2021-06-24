package Hot100

//101. 对称二叉树
//给定一个二叉树，检查它是否是镜像对称的。
//
//
//
//例如，二叉树 [1,2,2,3,4,4,3] 是对称的。
//
//1
/// \
//2   2
/// \ / \
//3  4 4  3
//
//
//但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:
//
//1
/// \
//2   2
//\   \
//3    3
//
//
//进阶：
//
//你可以运用递归和迭代两种方法解决这个问题吗？

// 迭代实现（逻辑队列）
func isSymmetric3(root *TreeNode) bool {
	if root == nil {
		return true
	}

	LeftQueue := []*TreeNode{root.Left}
	RightQueue := []*TreeNode{root.Right}

	for {
		LeftSize := len(LeftQueue)
		RightSize := len(RightQueue)

		if LeftSize != RightSize {
			return false
		}
		if LeftSize == 0 {
			break
		}

		for i := 0; i < LeftSize; i++ {
			LeftNode := LeftQueue[0]
			LeftQueue = LeftQueue[1:]
			RightNode := RightQueue[0]
			RightQueue = RightQueue[1:]
			if LeftNode == nil && RightNode == nil {
				continue
			}
			if LeftNode == nil && RightNode != nil || LeftNode != nil && RightNode == nil {
				return false
			}
			if LeftNode.Val != RightNode.Val {
				return false
			}
			LeftQueue = append(LeftQueue, LeftNode.Left)
			LeftQueue = append(LeftQueue, LeftNode.Right)
			RightQueue = append(RightQueue, RightNode.Right)
			RightQueue = append(RightQueue, RightNode.Left)
		}
	}

	return true
}

// 迭代实现（队列）
func isSymmetric2(root *TreeNode) bool {
	if root == nil {
		return true
	}

	LeftQueue := NewQueue()
	RightQueue := NewQueue()

	LeftQueue.Offer(root.Left)
	RightQueue.Offer(root.Right)

	for {
		LeftSize := LeftQueue.Size()
		RightSize := RightQueue.Size()

		if LeftSize != RightSize {
			return false
		}
		if LeftSize == 0 {
			break
		}

		for i := 0; i < LeftSize; i++ {
			LeftNode := LeftQueue.Poll().(*TreeNode)
			RightNode := RightQueue.Poll().(*TreeNode)
			if LeftNode == nil && RightNode == nil {
				continue
			}
			if LeftNode == nil && RightNode != nil || LeftNode != nil && RightNode == nil {
				return false
			}
			if LeftNode.Val != RightNode.Val {
				return false
			}
			LeftQueue.Offer(LeftNode.Left)
			LeftQueue.Offer(LeftNode.Right)
			RightQueue.Offer(RightNode.Right)
			RightQueue.Offer(RightNode.Left)
		}
	}

	return true
}

// 递归实现（代码优化）
func isSymmetric4(root *TreeNode) bool {
	return isSymmetricTree(root, root)
}

// 递归实现
func isSymmetric1(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSymmetricTree(root.Left, root.Right)
}

func isSymmetricTree(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	} else if root1 == nil || root2 == nil {
		return false
	}
	return root1.Val == root2.Val && isSymmetricTree(root1.Left, root2.Right) && isSymmetricTree(root1.Right, root2.Left)
}
