package Hot100

import "container/heap"

//23. 合并K个升序链表
//给你一个链表数组，每个链表都已经按升序排列。
//
//请你将所有链表合并到一个升序链表中，返回合并后的链表。
//
//
//
//示例 1：
//
//输入：lists = [[1,4,5],[1,3,4],[2,6]]
//输出：[1,1,2,3,4,4,5,6]
//解释：链表数组如下：
//[
//1->4->5,
//1->3->4,
//2->6
//]
//将它们合并到一个有序链表中得到。
//1->1->2->3->4->4->5->6
//示例 2：
//
//输入：lists = []
//输出：[]
//示例 3：
//
//输入：lists = [[]]
//输出：[]
//
//
//提示：
//
//k == lists.length
//0 <= k <= 10^4
//0 <= lists[i].length <= 500
//-10^4 <= lists[i][j] <= 10^4
//lists[i] 按 升序 排列
//lists[i].length 的总和不超过 10^4

type PriorityQueue []*ListNode

func (t *PriorityQueue) Len() int {
	return len(*t) //
}
func (t *PriorityQueue) Less(i, j int) bool {
	return (*t)[i].Val < (*t)[j].Val
}
func (t *PriorityQueue) Swap(i, j int) {
	(*t)[i], (*t)[j] = (*t)[j], (*t)[i]
}
func (t *PriorityQueue) Push(x interface{}) {
	*t = append(*t, x.(*ListNode))
}
func (t *PriorityQueue) Pop() interface{} {
	n := len(*t)
	x := (*t)[n-1]
	*t = (*t)[:n-1]
	return x
}
func (t *PriorityQueue) initQueue() {
	heap.Init(t)
}
func (t *PriorityQueue) isEmpty() bool {
	return t.Len() == 0
}
func (t *PriorityQueue) Poll() interface{} {
	return heap.Pop(t)
}
func (t *PriorityQueue) Offer(x interface{}) {
	heap.Push(t, x)
}

// 优先队列
func mergeKLists(lists []*ListNode) *ListNode {
	ret := new(ListNode)
	p := ret

	queue := PriorityQueue{}
	for id, list := range lists {
		if list != nil {
			queue = append(queue, list)
			lists[id] = lists[id].Next
		}
	}

	queue.initQueue()

	for {
		if queue.isEmpty() {
			break
		}
		p.Next = queue.Poll().(*ListNode)
		p = p.Next
		if p.Next != nil {
			queue.Offer(p.Next)
		}
	}

	return ret.Next
}

// 归并排序
func mergeKLists3(lists []*ListNode) *ListNode {
	return merge(lists, 0, len(lists)-1)
}

func merge(lists []*ListNode, left int, right int) *ListNode {
	if left == right {
		return lists[left]
	}

	if left > right {
		return nil
	}

	mid := (left + right) >> 1

	return mergeTwoLists(merge(lists, left, mid), merge(lists, mid+1, right))
}

// 改造多个链表之间的连接关系
func mergeKLists2(lists []*ListNode) *ListNode {
	ret := new(ListNode)
	p := ret

	for {
		minVal := 10001
		minId := -1
		for id, list := range lists {
			if list == nil {
				continue
			}
			if list.Val < minVal {
				minVal = list.Val
				minId = id
			}
		}
		if minId == -1 {
			break
		}
		p.Next = lists[minId]
		p = p.Next
		lists[minId] = lists[minId].Next
	}

	return ret.Next
}

// 新建一个新的链表
func mergeKLists1(lists []*ListNode) *ListNode {
	ret := new(ListNode)
	p := ret

	for {
		minVal := 10001
		minId := -1
		for id, list := range lists {
			if list == nil {
				continue
			}
			if list.Val < minVal {
				minVal = list.Val
				minId = id
			}
		}
		if minId == -1 {
			break
		}
		p.Next = &ListNode{Val: minVal}
		p = p.Next
		lists[minId] = lists[minId].Next
	}

	return ret.Next
}
