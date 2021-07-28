package Hot100

//206. 反转链表
//给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
//
//
//示例 1：
//
//
//输入：head = [1,2,3,4,5]
//输出：[5,4,3,2,1]
//示例 2：
//
//
//输入：head = [1,2]
//输出：[2,1]
//示例 3：
//
//输入：head = []
//输出：[]
//
//
//提示：
//
//链表中节点的数目范围是 [0, 5000]
//-5000 <= Node.val <= 5000
//
//
//进阶：链表可以选用迭代或递归方式完成反转。你能否用两种方法解决这道题？

// 迭代（更改连接方式）
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	p, q := head, head.Next
	head.Next = nil
	for q.Next != nil {
		r := q.Next
		q.Next = p
		p, q = q, r
	}
	q.Next = p

	return q
}

// 迭代（栈）
func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	p := head
	stack := []*ListNode{}
	for p != nil {
		stack = append(stack, p)
		p = p.Next
	}

	head = stack[len(stack)-1]
	p = head
	for {
		stack = stack[:len(stack)-1]
		if len(stack) <= 0 {
			break
		}
		p.Next = stack[len(stack)-1]
		p = p.Next
	}
	p.Next = nil

	return head
}

// 递归
func reverseList1(head *ListNode) *ListNode {
	return reverseListNode(head)
}
