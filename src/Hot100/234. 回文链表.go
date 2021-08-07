package Hot100

//234. 回文链表
//请判断一个链表是否为回文链表。
//
//示例 1:
//
//输入: 1->2
//输出: false
//示例 2:
//
//输入: 1->2->2->1
//输出: true
//进阶：
//你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？

// 快慢指针 + 反转链表
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	p, q := head, head
	for q != nil && q.Next != nil {
		p = p.Next
		q = q.Next.Next
	}

	newHead := reverseListNode(p)

	for newHead != nil {
		if newHead.Val != head.Val {
			return false
		}
		newHead = newHead.Next
		head = head.Next
	}
	return true
}

// 递归实现，超出时间限制（84 / 85）
func isPalindrome1(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	p, q := head, head.Next
	for q != nil && q.Next != nil {
		p = p.Next
		q = q.Next
	}
	if head.Val != q.Val {
		return false
	}
	p.Next = nil
	return isPalindrome(head.Next)
}
