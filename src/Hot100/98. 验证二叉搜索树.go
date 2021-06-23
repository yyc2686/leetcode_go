package Hot100

//98. 验证二叉搜索树
//给定一个二叉树，判断其是否是一个有效的二叉搜索树。
//
//假设一个二叉搜索树具有如下特征：
//
//节点的左子树只包含小于当前节点的数。
//节点的右子树只包含大于当前节点的数。
//所有左子树和右子树自身必须也是二叉搜索树。
//示例 1:
//
//输入:
//2
/// \
//1   3
//输出: true
//示例 2:
//
//输入:
//5
/// \
//1   4
/// \
//3   6
//输出: false
//解释: 输入为: [5,1,4,null,null,3,6]。
//根节点的值为 5 ，但是其右子节点值为 4 。

const INT_MAX = int(^uint(0) >> 1)
const INT_MIN = ^INT_MAX

// 递归实现
func isValidBST(root *TreeNode) bool {
	// 空树与孤立节点为BST
	if root == nil || root.Left == nil && root.Right == nil {
		return true
	}

	// 判断左分支最大值 < 根节点值 < 右分支最小值
	if (root.Left != nil && maxValBST(root.Left) >= root.Val) || (root.Right != nil && root.Val >= minValBST(root.Right)) {
		return false
	}

	// 保证左右孩子都是BST
	if root.Left != nil {
		if root.Left.Val >= root.Val || !isValidBST(root.Left) {
			return false
		}
	}
	if root.Right != nil {
		if root.Right.Val <= root.Val || !isValidBST(root.Right) {
			return false
		}
	}

	return true
}

func minValBST(root *TreeNode) int {
	if root.Left != nil {
		return minValBST(root.Left)
	} else {
		return root.Val
	}
}

func maxValBST(root *TreeNode) int {
	if root.Right != nil {
		return maxValBST(root.Right)
	} else {
		return root.Val
	}
}
