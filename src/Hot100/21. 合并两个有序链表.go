package Hot100

//21. 合并两个有序链表
//将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
//
//
//
//示例 1：
//
//
//输入：l1 = [1,2,4], l2 = [1,3,4]
//输出：[1,1,2,3,4,4]
//示例 2：
//
//输入：l1 = [], l2 = []
//输出：[]
//示例 3：
//
//输入：l1 = [], l2 = [0]
//输出：[0]
//
//
//提示：
//
//两个链表的节点数目范围是 [0, 50]
//-100 <= Node.val <= 100
//l1 和 l2 均按 非递减顺序 排列

// 递归
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}

	if l1.Val <= l2.Val {
		ret := l1
		ret.Next = mergeTwoLists(l1.Next, l2)
		return ret
	} else {
		ret := l2
		ret.Next = mergeTwoLists(l1, l2.Next)
		return ret
	}
}

// 直接改造l1和l2上的元素链接关系（优化，更改链接方式后不需要置空）
func mergeTwoLists3(l1 *ListNode, l2 *ListNode) *ListNode {
	ret := new(ListNode)
	p := ret
	for {
		if l1 == nil || l2 == nil {
			break
		}
		if l1.Val <= l2.Val {
			p.Next = l1
			l1 = l1.Next
		} else {
			p.Next = l2
			l2 = l2.Next
		}
		p = p.Next
	}
	if l2 != nil {
		p.Next = l2
	} else if l1 != nil {
		p.Next = l1
	}
	return ret.Next
}

// 直接改造l1和l2上的元素链接关系
func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	ret := new(ListNode)
	r := ret
	for {
		if l1 == nil || l2 == nil {
			break
		}
		p := l1.Next
		q := l2.Next

		if l1.Val <= l2.Val {
			r.Next = l1
			l1.Next = nil
			l1 = p
		} else {
			r.Next = l2
			l2.Next = nil
			l2 = q
		}
		r = r.Next
	}
	if l2 != nil {
		r.Next = l2
	} else if l1 != nil {
		r.Next = l1
	}
	return ret.Next
}

// 新建一个新的链表
func mergeTwoLists1(l1 *ListNode, l2 *ListNode) *ListNode {
	ret := new(ListNode)
	p := ret
	for {
		if l1 == nil || l2 == nil {
			break
		}
		node := new(ListNode)
		if l1.Val <= l2.Val {
			node.Val = l1.Val
			l1 = l1.Next
		} else {
			node.Val = l2.Val
			l2 = l2.Next
		}
		p.Next = node
		p = p.Next
	}
	if l2 != nil {
		p.Next = l2
	} else if l1 != nil {
		p.Next = l1
	}
	return ret.Next
}
