package Hot100

//148. 排序链表
//给你链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。
//
//进阶：
//
//你可以在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序吗？
//
//
//示例 1：
//
//
//输入：head = [4,2,1,3]
//输出：[1,2,3,4]
//示例 2：
//
//
//输入：head = [-1,5,3,4,0]
//输出：[-1,0,3,4,5]
//示例 3：
//
//输入：head = []
//输出：[]
//
//
//提示：
//
//链表中节点的数目在范围 [0, 5 * 104] 内
//-105 <= Node.val <= 105

// 递归实现+归并排序（时间复杂度O(nlogn)，空间复杂度O(1)）
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	left, right := middleNode(head)
	return mergeTwoLists(sortList(left), sortList(right))
}

// 找到链表中点并断开（快慢指针）
func middleNode(head *ListNode) (*ListNode, *ListNode) {
	if head == nil {
		return head, head
	}

	fast, slow := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	fast = slow.Next
	slow.Next = nil
	return head, fast
}

// 递归实现+冒泡排序？（时间复杂度O(n^2)）
func sortList1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := sortList(head.Next)
	// 插入头
	if head.Val <= newHead.Val {
		head.Next = newHead
		return head
	}

	// 插入中间和尾部
	p := newHead
	for p.Next != nil && p.Next.Val < head.Val {
		p = p.Next
	}
	head.Next = p.Next
	p.Next = head
	return newHead

}
