package Hot100

//2. 两数相加
//给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
//
//请你将两个数相加，并以相同形式返回一个表示和的链表。
//
//你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
//
//
//
//示例 1：
//
//
//输入：l1 = [2,4,3], l2 = [5,6,4]
//输出：[7,0,8]
//解释：342 + 465 = 807.
//示例 2：
//
//输入：l1 = [0], l2 = [0]
//输出：[0]
//示例 3：
//
//输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
//输出：[8,9,9,9,0,0,0,1]
//
//
//提示：
//
//每个链表中的节点数在范围 [1, 100] 内
//0 <= Node.val <= 9
//题目数据保证列表表示的数字不含前导零

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func ConvertSliceToListNode(slice []int) *ListNode {
	if len(slice) == 0 {
		return nil
	}

	head := ListNode{-1, nil}
	p := &head
	for _, v := range slice {
		tmp := ListNode{v, nil}
		p.Next = &tmp
		p = p.Next
	}
	return head.Next
}

func reverseListNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := reverseListNode(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	jin := 0

	//l1 = reverseListNode(l1)
	//l2 = reverseListNode(l2)

	head := ListNode{-1, nil}
	p := &head

	//head := &ListNode{-1, nil}
	//p := head

	for {
		if l1 == nil && l2 == nil && jin == 0 {
			break
		}

		if l1 == nil {
			l1 = &ListNode{0, nil}
		}
		if l2 == nil {
			l2 = &ListNode{0, nil}
		}

		val := l1.Val + l2.Val + jin
		if val >= 10 {
			jin = 1
			val -= 10
		} else {
			jin = 0
		}

		node := &ListNode{val, nil}
		p.Next = node
		p = p.Next
		l1 = l1.Next
		l2 = l2.Next
	}

	return head.Next
}
