package Hot100

//32. 最长有效括号
//给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。
//
//
//
//示例 1：
//
//输入：s = "(()"
//输出：2
//解释：最长有效括号子串是 "()"
//示例 2：
//
//输入：s = ")()())"
//输出：4
//解释：最长有效括号子串是 "()()"
//示例 3：
//
//输入：s = ""
//输出：0
//
//
//提示：
//
//0 <= s.length <= 3 * 104
//s[i] 为 '(' 或 ')'

type Stack interface {
	Push(e Element) //向队列中添加元素
	Pop() Element   //移除队列中最前面的元素
	Peek() Element  //移除队列中最前面的元素
	Clear() bool    //清空队列
	Size() int      //获取队列的元素个数
	IsEmpty() bool  //判断队列是否是空
}

func NewStack() *sliceEntry {
	return &sliceEntry{}
}

//向队列中添加元素
func (entry *sliceEntry) Push(e Element) {
	entry.element = append(entry.element, e)
}

//移除队列中最前面的额元素
func (entry *sliceEntry) Pop() Element {
	if entry.IsEmpty() {
		return nil
	}

	size := entry.Size()
	ret := entry.element[size-1]
	entry.element = entry.element[:size-1]
	return ret
}

// 栈顶元素
func (entry *sliceEntry) Peek() Element {
	//判断栈是否空
	if entry.IsEmpty() {
		return nil
	}

	return entry.element[entry.Size()-1]
}

// 使用(类似)栈优化，使用top记录栈顶位置
func longestValidParentheses4(s string) int {
	//算法核心：
	//    始终保持栈底元素为当前已经遍历过的元素中「最后一个没有被匹配的右括号的下标」，
	//    栈里其他元素维护左括号的下标
	//
	//1. 对于遇到的每个 \text{‘(’}‘(’ ，我们将它的下标放入栈中
	//2. 对于遇到的每个 \text{‘)’}‘)’ ，我们先弹出栈顶元素表示匹配了当前右括号：
	//		(1) 如果栈为空，说明当前的右括号为没有被匹配的右括号，我们将其下标放入栈中来更新我们之前提到的「最后一个没有被匹配的右括号的下标」
	//		(2) 如果栈不为空，当前右括号的下标减去栈顶元素即为「以该右括号为结尾的最长有效括号的长度」
	//Note:
	// 		如果一开始栈为空，第一个字符为左括号的时候我们会将其放入栈中，这样就不满足提及的「最后一个没有被匹配的右括号的下标」，
	// 		为了保持统一，我们在一开始的时候往栈中放入一个值为 -1−1 的元素。

	size := len(s)
	ret := 0

	if size <= 1 {
		return 0
	}

	stack := []int{-1}
	top := 0

	for i := 0; i < size; i++ {
		if s[i] == '(' {
			top++
			stack = append(stack, i)
		} else {
			stack = stack[:top]
			top--
			if top == -1 {
				top++
				stack = append(stack, i)
			} else {
				ret = maxInt(ret, i-stack[top])
			}
		}
	}

	return ret
}

// 使用(类似)栈优化
func longestValidParentheses3(s string) int {
	//算法核心：
	//    始终保持栈底元素为当前已经遍历过的元素中「最后一个没有被匹配的右括号的下标」，
	//    栈里其他元素维护左括号的下标
	//
	//1. 对于遇到的每个 \text{‘(’}‘(’ ，我们将它的下标放入栈中
	//2. 对于遇到的每个 \text{‘)’}‘)’ ，我们先弹出栈顶元素表示匹配了当前右括号：
	//		(1) 如果栈为空，说明当前的右括号为没有被匹配的右括号，我们将其下标放入栈中来更新我们之前提到的「最后一个没有被匹配的右括号的下标」
	//		(2) 如果栈不为空，当前右括号的下标减去栈顶元素即为「以该右括号为结尾的最长有效括号的长度」
	//Note:
	// 		如果一开始栈为空，第一个字符为左括号的时候我们会将其放入栈中，这样就不满足提及的「最后一个没有被匹配的右括号的下标」，
	// 		为了保持统一，我们在一开始的时候往栈中放入一个值为 -1−1 的元素。

	size := len(s)
	ret := 0

	if size <= 1 {
		return 0
	}

	stack := []int{-1}

	for i := 0; i < size; i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				ret = maxInt(ret, i-stack[len(stack)-1])
			}
		}
	}

	return ret
}

// 使用栈实现，保持栈底元素为当前已经遍历过的元素中「最后一个没有被匹配的右括号的下标」
func longestValidParentheses2(s string) int {
	//算法核心：
	//    始终保持栈底元素为当前已经遍历过的元素中「最后一个没有被匹配的右括号的下标」，
	//    栈里其他元素维护左括号的下标
	//
	//1. 对于遇到的每个 \text{‘(’}‘(’ ，我们将它的下标放入栈中
	//2. 对于遇到的每个 \text{‘)’}‘)’ ，我们先弹出栈顶元素表示匹配了当前右括号：
	//		(1) 如果栈为空，说明当前的右括号为没有被匹配的右括号，我们将其下标放入栈中来更新我们之前提到的「最后一个没有被匹配的右括号的下标」
	//		(2) 如果栈不为空，当前右括号的下标减去栈顶元素即为「以该右括号为结尾的最长有效括号的长度」
	//Note:
	// 		如果一开始栈为空，第一个字符为左括号的时候我们会将其放入栈中，这样就不满足提及的「最后一个没有被匹配的右括号的下标」，
	// 		为了保持统一，我们在一开始的时候往栈中放入一个值为 -1−1 的元素。

	size := len(s)
	ret := 0

	if size <= 1 {
		return 0
	}

	stack := NewStack()
	stack.Push(-1)

	for i := 0; i < size; i++ {
		if s[i] == '(' {
			stack.Push(i)
		} else {
			stack.Pop()
			if stack.IsEmpty() {
				stack.Push(i)
			} else {
				ret = maxInt(ret, i-stack.Peek().(int))
			}
		}
	}

	return ret
}

// 使用动态规划实现，dp[i]：以第i个字符结尾的最长有效括号长度
func longestValidParentheses1(s string) int {
	// 1. s[i] == '(' dp[i] = 0
	// 2. s[i] == ')'
	//	  (1) s[i-1] == '(', dp[i] = dp[i-2]+2
	//	  (2) s[i-1] == ')',
	//			if s[i-1-dp[i-1]] == '(': dp[i] = dp[i-1] + 2 + dp[i-1-dp[i-1]-1]

	size := len(s)
	ret := 0

	if size <= 1 {
		return 0
	}

	dp := make([]int, size)
	dp[0] = 0
	if s[0] == '(' && s[1] == ')' {
		dp[1] = 2
		ret = 2
	} else {
		dp[1] = 0
	}

	for i := 2; i < size; i++ {
		if s[i] == '(' {
			dp[i] = 0
		} else if s[i-1] == '(' {
			dp[i] = dp[i-2] + 2
		} else if i-1-dp[i-1] >= 0 && s[i-1-dp[i-1]] == '(' {
			if i-1-dp[i-1]-1 >= 0 {
				dp[i] = dp[i-1] + 2 + dp[i-1-dp[i-1]-1]
			} else {
				dp[i] = dp[i-1] + 2
			}
		}
		ret = maxInt(ret, dp[i])
	}

	return ret
}

//有效括号性质：左括号数>=右括号数，两趟遍历（失败，原因在于必要条件无法精准描述括号的连续性）
// !!! 其实可行，更新ret时刻为left==right !!!
func longestValidParentheses(s string) int {
	left, right := 0, 0
	size := len(s)
	ret := 0

	// 从左往右扫描
	for i := 0; i < size; i++ {
		if s[i] == '(' {
			left++
		} else {
			right++
		}

		if left < right {
			left, right = 0, 0
		} else if left == right {
			ret = maxInt(ret, left)
		}
	}

	left, right = 0, 0
	// 从右往左扫描
	for i := size - 1; i >= 0; i-- {
		if s[i] == '(' {
			left++
		} else {
			right++
		}

		if right < left {
			left, right = 0, 0
		} else if left == right {
			ret = maxInt(ret, right)
		}
	}

	return 2 * ret
}
