package Hot100

//105. 从前序与中序遍历序列构造二叉树
//根据一棵树的前序遍历与中序遍历构造二叉树。
//
//注意:
//你可以假设树中没有重复的元素。
//
//例如，给出
//
//前序遍历 preorder = [3,9,20,15,7]
//中序遍历 inorder = [9,3,15,20,7]
//返回如下的二叉树：
//
//3
/// \
//9  20
///  \
//15   7

// 非递归实现（栈/逻辑栈）
func buildTree(preorder []int, inorder []int) *TreeNode {
	// 综上所述，我们用一个栈保存已经遍历过的节点，
	// 遍历前序遍历的数组，一直作为当前根节点的左子树，直到当前节点和中序遍历的数组的节点相等了，
	// 那么我们正序遍历中序遍历的数组，倒着遍历已经遍历过的根节点（用栈的 pop 实现），
	// 找到最后一次相等的位置，把它作为该节点的右子树。

	size := len(preorder)
	if size == 0 {
		return nil
	}

	pre, in := 0, 0

	root := &TreeNode{preorder[pre], nil, nil}
	curRoot := root
	stack := []*TreeNode{root}

	for pre++; pre < size; pre++ {
		// 出现了当前节点的值和中序遍历数组的值相等，寻找的是谁的右子树
		if curRoot.Val == inorder[in] {
			for ; len(stack) != 0 && stack[len(stack)-1].Val == inorder[in]; in++ {
				curRoot = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
			}
			curRoot.Right = &TreeNode{preorder[pre], nil, nil}
			curRoot = curRoot.Right
		} else {
			// 否则一直作为左子树
			curRoot.Left = &TreeNode{preorder[pre], nil, nil}
			curRoot = curRoot.Left
		}
		stack = append(stack, curRoot)
	}
	return root
}

// 递归实现
func buildTree1(preorder []int, inorder []int) *TreeNode {
	size := len(preorder)
	if size == 0 {
		return nil
	}

	id := 0
	for ; id < size; id++ {
		if inorder[id] == preorder[0] {
			break
		}
	}
	root := new(TreeNode)
	root.Val = preorder[0]
	root.Left = buildTree(preorder[1:id+1], inorder[:id])
	root.Right = buildTree(preorder[id+1:], inorder[id+1:])

	return root
}
