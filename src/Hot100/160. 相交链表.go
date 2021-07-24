package Hot100

//160. 相交链表
//给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。如果两个链表没有交点，返回 null 。
//
//图示两个链表在节点 c1 开始相交：
//
//
//
//题目数据 保证 整个链式结构中不存在环。
//
//注意，函数返回结果后，链表必须 保持其原始结构 。
//
//
//
//示例 1：
//
//
//
//输入：intersectVal = 8, listA = [4,1,8,4,5], listB = [5,0,1,8,4,5], skipA = 2, skipB = 3
//输出：Intersected at '8'
//解释：相交节点的值为 8 （注意，如果两个链表相交则不能为 0）。
//从各自的表头开始算起，链表 A 为 [4,1,8,4,5]，链表 B 为 [5,0,1,8,4,5]。
//在 A 中，相交节点前有 2 个节点；在 B 中，相交节点前有 3 个节点。
//示例 2：
//
//
//
//输入：intersectVal = 2, listA = [0,9,1,2,4], listB = [3,2,4], skipA = 3, skipB = 1
//输出：Intersected at '2'
//解释：相交节点的值为 2 （注意，如果两个链表相交则不能为 0）。
//从各自的表头开始算起，链表 A 为 [0,9,1,2,4]，链表 B 为 [3,2,4]。
//在 A 中，相交节点前有 3 个节点；在 B 中，相交节点前有 1 个节点。
//示例 3：
//
//
//
//输入：intersectVal = 0, listA = [2,6,4], listB = [1,5], skipA = 3, skipB = 2
//输出：null
//解释：从各自的表头开始算起，链表 A 为 [2,6,4]，链表 B 为 [1,5]。
//由于这两个链表不相交，所以 intersectVal 必须为 0，而 skipA 和 skipB 可以是任意值。
//这两个链表不相交，因此返回 null 。
//
//
//提示：
//
//listA 中节点数目为 m
//listB 中节点数目为 n
//0 <= m, n <= 3 * 104
//1 <= Node.val <= 105
//0 <= skipA <= m
//0 <= skipB <= n
//如果 listA 和 listB 没有交点，intersectVal 为 0
//如果 listA 和 listB 有交点，intersectVal == listA[skipA + 1] == listB[skipB + 1]
//
//
//进阶：你能否设计一个时间复杂度 O(n) 、仅用 O(1) 内存的解决方案？

// 反转链表，找到第一个不相交的节点（改变连接关系，混乱！）
//func getIntersectionNode(headA, headB *ListNode) *ListNode {
//	if headA == nil || headB == nil {
//		return nil
//	}
//	headA = reverseListNode(headA)
//	headB = reverseListNode(headB)
//	p := headA
//	for ; p != nil && headB != nil && p == headB; {
//		p = p.Next
//		headB = headB.Next
//	}
//	if p!=nil{
//		p.Next = nil
//	}
//	return reverseListNode(headA)
//}

// 双指针（代码优化）（有且仅有两种结果：不相交，二者都为nil；相加，找到相交点，均可退出循环）
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pa, pb := headA, headB
	for pa != pb {
		if pa == nil {
			pa = headB
		} else {
			pa = pa.Next
		}

		if pb == nil {
			pb = headA
		} else {
			pb = pb.Next
		}
	}

	return pa
}

// 双指针（相交的两链表必有相同的一段）（关键在于消除长度差）
func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pa, pb := headA, headB

	// 消除长度差
	for pa != nil && pb != nil {
		pa = pa.Next
		pb = pb.Next
	}

	if pa == nil {
		for pb != nil {
			headB = headB.Next
			pb = pb.Next
		}
	}
	if pb == nil {
		for pa != nil {
			headA = headA.Next
			pa = pa.Next
		}
	}

	// 寻找相交点
	for headA != nil && headB != nil {
		if headA == headB {
			return headA
		}
		headA = headA.Next
		headB = headB.Next
	}
	return nil
}

// 哈希表法（记录每个节点）
func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	m := make(map[*ListNode]bool, 0)
	for headA != nil {
		m[headA] = true
		headA = headA.Next
	}

	for headB != nil {
		if m[headB] {
			return headB
		}
		headB = headB.Next
	}
	return nil
}
