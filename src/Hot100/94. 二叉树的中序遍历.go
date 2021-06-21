package Hot100

//94. 二叉树的中序遍历
//给定一个二叉树的根节点 root ，返回它的 中序 遍历。
//
//
//
//示例 1：
//
//
//输入：root = [1,null,2,3]
//输出：[1,3,2]
//示例 2：
//
//输入：root = []
//输出：[]
//示例 3：
//
//输入：root = [1]
//输出：[1]
//示例 4：
//
//
//输入：root = [1,2]
//输出：[2,1]
//示例 5：
//
//
//输入：root = [1,null,2]
//输出：[1,2]
//
//
//提示：
//
//树中节点数目在范围 [0, 100] 内
//-100 <= Node.val <= 100
//
//
//进阶: 递归算法很简单，你可以通过迭代算法完成吗？

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func ConvertSliceToTreeNode(slice []int) *TreeNode {

	size := len(slice)
	if size == 0 {
		return nil
	}

	root := &TreeNode{slice[0], nil, nil}

	queue := NewQueue()
	queue.Offer(root)

	k := 1

	for {
		if k >= size || queue.IsEmpty() {
			break
		}
		cap := queue.Size()
		for i := 0; i < cap; i++ {
			node := queue.Pop().(*TreeNode)
			if k < size {
				if slice[k] != -1 {
					node.Left = &TreeNode{slice[k], nil, nil}
					queue.Offer(node.Left)
				}
				k++
			}
			if k < size {
				if slice[k] != -1 {
					node.Right = &TreeNode{slice[k], nil, nil}
					queue.Offer(node.Right)
				}
				k++
			}
		}

	}

	return root
}

// 非递归实现（逻辑栈）（2）双重循环
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	ret := make([]int, 0)

	stack := make([]*TreeNode, 0)
	for {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		if len(stack) == 0 {
			break
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ret = append(ret, root.Val)
		root = root.Right

	}

	return ret
}

// 非递归实现（栈）（2）双重循环
func inorderTraversal4(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	ret := make([]int, 0)

	stack := NewStack()
	for {
		for root != nil {
			stack.Push(root)
			root = root.Left
		}
		if stack.IsEmpty() {
			break
		}
		root = stack.Pop().(*TreeNode)
		ret = append(ret, root.Val)
		root = root.Right

	}

	return ret
}

// 非递归实现（逻辑栈）（1）一路向左
func inorderTraversal3(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	ret := make([]int, 0)

	stack := []*TreeNode{root}

	for {
		if root.Left != nil {
			stack = append(stack, root.Left)
			root = root.Left
		} else {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			root.Left = nil
			ret = append(ret, root.Val)
			if root.Right != nil {
				stack = append(stack, root.Right)
				root = root.Right
			}
		}
		if len(stack) == 0 {
			break
		}
	}
	return ret
}

// 非递归实现（栈）（1）一路向左
func inorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	ret := make([]int, 0)

	stack := NewStack()
	stack.Push(root)

	for {
		if root.Left != nil {
			stack.Push(root.Left)
			root = root.Left
		} else {
			root = stack.Pop().(*TreeNode)
			root.Left = nil
			ret = append(ret, root.Val)
			if root.Right != nil {
				stack.Push(root.Right)
				root = root.Right
			}
		}
		if stack.IsEmpty() {
			break
		}
	}
	return ret
}

// 递归实现
func inorderTraversal1(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	return append(append(inorderTraversal(root.Left), root.Val), inorderTraversal(root.Right)...)
}
