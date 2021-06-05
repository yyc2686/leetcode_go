package Hot100

//19. 删除链表的倒数第 N 个结点
//给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
//
//进阶：你能尝试使用一趟扫描实现吗？
//
//
//
//示例 1：
//
//
//输入：head = [1,2,3,4,5], n = 2
//输出：[1,2,3,5]
//示例 2：
//
//输入：head = [1], n = 1
//输出：[]
//示例 3：
//
//输入：head = [1,2], n = 1
//输出：[1]
//
//
//提示：
//
//链表中结点的数目为 sz
//1 <= sz <= 30
//0 <= Node.val <= 100
//1 <= n <= sz

// 双指针
func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	left := head
	right := head
	for i := 0; i < n; i++ {
		right = right.Next
	}

	for {
		if right == nil {
			head = head.Next
			return head
		} else if right.Next != nil {
			left = left.Next
			right = right.Next
		} else {
			break
		}
	}

	left.Next = left.Next.Next

	return head

}

// 两次反转链表
func reverseLinkList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := reverseLinkList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	newHead := reverseLinkList(head)
	p := newHead

	if n == 1 {
		return reverseLinkList(newHead.Next)
	}

	for i := 0; i < n-2; i++ {
		p = p.Next
	}
	p.Next = p.Next.Next
	return reverseLinkList(newHead)

}

// 辅助数+递归
var cur int

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	cur = 0
	return helperRemoveNthFromEnd(head, n)
}

func helperRemoveNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}

	head.Next = helperRemoveNthFromEnd(head.Next, n)
	cur++
	if cur == n {
		return head.Next
	} else {
		return head
	}

}
