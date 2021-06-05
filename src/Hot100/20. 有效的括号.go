package Hot100

//20. 有效的括号
//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
//
//有效字符串需满足：
//
//左括号必须用相同类型的右括号闭合。
//左括号必须以正确的顺序闭合。
//
//
//示例 1：
//
//输入：s = "()"
//输出：true
//示例 2：
//
//输入：s = "()[]{}"
//输出：true
//示例 3：
//
//输入：s = "(]"
//输出：false
//示例 4：
//
//输入：s = "([)]"
//输出：false
//示例 5：
//
//输入：s = "{[]}"
//输出：true
//
//
//提示：
//
//1 <= s.length <= 104
//s 仅由括号 '()[]{}' 组成

var parentheses = map[byte]byte{
	'(': ')',
	'[': ']',
	'{': '}',
}

// 栈

//使用数组来模拟一个栈的使用

//type Stack interface {
//	Push(e Element) // 向栈中添加元素
//	Pop() Element   // 弹出栈顶元素
//	Peek() Element  // 显示栈顶元素
//	Size() int      // 获取栈的元素个数
//	IsEmpty() bool  // 判断队列是否为空
//}
//
//func NewStack() *sliceEntry {
//	return &sliceEntry{}
//}
//
////入栈
//func (entry *sliceEntry) Push(e Element) {
//	//放入数据
//	entry.element = append([]Element{e}, entry.element...)
//}
//
////出栈
//func (entry *sliceEntry) Pop() Element {
//	//判断栈是否空
//	if entry.IsEmpty() {
//		return nil
//	}
//
//	firstElement := entry.element[0]
//	entry.element = entry.element[1:]
//	return firstElement
//}
//
//// 栈顶元素
//func (entry *sliceEntry) Peek() Element {
//	//判断栈是否空
//	if entry.IsEmpty() {
//		return nil
//	}
//	return entry.element[0]
//}

func isValid(s string) bool {
	stack := NewStack()
	size := len(s)
	for i := 0; i < size; i++ {
		if _, ok := parentheses[s[i]]; ok {
			stack.Push(s[i])
		} else {
			tmp := stack.Pop()
			if tmp == nil || s[i] != parentheses[tmp.(byte)] {
				return false
			}

		}
	}
	return stack.IsEmpty()

}

// 左右扫描计数（不适用于多种括号情形）
//func isValid(s string) bool {
//	left, right := 0, 0
//	for i := 0; i < len(s); i++ {
//		if _, ok := parentheses[s[i]]; ok {
//			left++
//		} else {
//			right++
//		}
//		if left < right {
//			return false
//		}
//	}
//	if left != right {
//		return false
//	}
//
//	left, right = 0, 0
//	for i := len(s) - 1; i >= 0; i-- {
//		if _, ok := parentheses[s[i]]; ok {
//			left++
//		} else {
//			right++
//		}
//		if right < left {
//			return false
//		}
//	}
//
//	return left == right
//
//}
//	if size%2 == 1 {
//		return false

// 双指针（不可行）
//func isValid(s string) bool {
//	size := len(s)
//	left, right := 0, size-1
//	}
//
//	for {
//		if left >= right {
//			break
//		}
//
//		if s[left] != s[right] {
//			return false
//		}
//
//		left++
//		right--
//	}
//
//	return true
//}
