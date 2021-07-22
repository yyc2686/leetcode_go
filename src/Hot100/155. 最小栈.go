package Hot100

//155. 最小栈
//设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。
//
//push(x) —— 将元素 x 推入栈中。
//pop() —— 删除栈顶的元素。
//top() —— 获取栈顶元素。
//getMin() —— 检索栈中的最小元素。
//
//
//示例:
//
//输入：
//["MinStack","push","push","push","getMin","pop","top","getMin"]
//[[],[-2],[0],[-3],[],[],[],[]]
//
//输出：
//[null,null,null,null,-3,null,0,-2]
//
//解释：
//MinStack minStack = new MinStack();
//minStack.push(-2);
//minStack.push(0);
//minStack.push(-3);
//minStack.getMin();   --> 返回 -3.
//minStack.pop();
//minStack.top();      --> 返回 0.
//minStack.getMin();   --> 返回 -2.
//
//
//提示：
//
//pop、top 和 getMin 操作总是在 非空栈 上调用。

type MinStack struct {
	stack         []int
	minIndexStack []int
}

/** initialize your data structure here. */
func MinStackConstructor() MinStack {
	minStack := MinStack{}
	minStack.stack = make([]int, 0)
	minStack.minIndexStack = make([]int, 0)
	return minStack
}

// 空间优化，只有与当前top minIndex不同时才出入栈
func (this *MinStack) Push(val int) {
	size := len(this.stack)
	if size == 0 {
		this.stack = []int{val}
		this.minIndexStack = []int{0}
	} else {
		this.stack = append(this.stack, val)
		if val <= this.stack[this.minIndexStack[len(this.minIndexStack)-1]] {
			this.minIndexStack = append(this.minIndexStack, size)
		}
	}
}

func (this *MinStack) Pop() {
	size := len(this.stack)
	if this.stack != nil && size != 0 {
		this.stack = this.stack[:size-1]
		if size-1 == this.minIndexStack[len(this.minIndexStack)-1] {
			this.minIndexStack = this.minIndexStack[:len(this.minIndexStack)-1]
		}
	}
}

// 空间优化前
func (this *MinStack) Push1(val int) {
	size := len(this.stack)
	if size == 0 {
		this.stack = []int{val}
		this.minIndexStack = []int{0}
	} else {
		this.stack = append(this.stack, val)
		if val <= this.stack[this.minIndexStack[size-1]] {
			this.minIndexStack = append(this.minIndexStack, size)
		} else {
			this.minIndexStack = append(this.minIndexStack, this.minIndexStack[size-1])
		}
	}
}

func (this *MinStack) Pop1() {
	size := len(this.stack)
	if this.stack != nil && size != 0 {
		this.stack = this.stack[:size-1]
		this.minIndexStack = this.minIndexStack[:size-1]
	}
}

func (this *MinStack) Top() int {
	size := len(this.stack)
	if this.stack != nil && size != 0 {
		return this.stack[size-1]
	}
	return -1
}

func (this *MinStack) GetMin() int {
	size := len(this.minIndexStack)
	if this.stack != nil && size != 0 {
		return this.stack[this.minIndexStack[size-1]]
	}
	return -1
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
