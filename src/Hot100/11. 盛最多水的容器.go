package Hot100

//11. 盛最多水的容器
//给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0) 。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
//
//说明：你不能倾斜容器。
//
//
//
//示例 1：
//
//
//
//输入：[1,8,6,2,5,4,8,3,7]
//输出：49
//解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。
//示例 2：
//
//输入：height = [1,1]
//输出：1
//示例 3：
//
//输入：height = [4,3,2,1,4]
//输出：16
//示例 4：
//
//输入：height = [1,2,1]
//输出：2
//
//
//提示：
//
//n = height.length
//2 <= n <= 3 * 104
//0 <= height[i] <= 3 * 104

////使用数组来模拟一个栈的使用
//type Stack struct {
//	MaxTop int   // 表示我们栈最大可以存放数个数
//	Top    int   // 表示栈顶, 因为栈顶固定，因此我们直接使用Top
//	arr    []int // 数组模拟栈
//}
//
////入栈
//func (this *Stack) Push(val int) (err error) {
//
//	//先判断栈是否满了
//	if this.Top == this.MaxTop-1 {
//		return errors.New("stack full")
//	}
//	this.Top++
//	//放入数据
//	this.arr[this.Top] = val
//
//	return
//}
//
////出栈
//func (this *Stack) Pop() (val int, err error) {
//	//判断栈是否空
//	if this.Top == -1 {
//		return -1, errors.New("stack empty")
//	}
//
//	//先取值，再 this.Top--
//	val = this.arr[this.Top]
//	this.Top--
//	return val, nil
//}
//
//// 栈顶元素
//func (this *Stack) Peek() (val int, err error) {
//	//判断栈是否空
//	if this.Top == -1 {
//		return -1, errors.New("stack empty")
//	}
//
//	//先取值，再 this.Top--
//	val = this.arr[this.Top]
//	return this.arr[this.Top], nil
//}

func maxArea(height []int) int {
	length := len(height)
	left, right := 0, length-1

	ret := 0
	for {
		ret = maxInt(ret, minInt(height[left], height[right])*(right-left))
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
		if left >= right {
			break
		}
	}
	return ret
}

////使用结构体管理栈
//type Stack struct {
//	arr       []int //切片
//	stackSize int   //栈中元素的个数
//}
//
////创建栈
//func newStack() *Stack {
//	return &Stack{arr: make([]int, 0)}
//}
//
////判断是否为空
//func (p *Stack) isEmpty() bool {
//	return p.stackSize == 0
//}
//
////栈的大小
//func (p *Stack) size() int {
//	return p.stackSize
//}
//
////返回栈顶元素
//func (p *Stack) peek() (int, error) {
//	if p.isEmpty() { //说明栈为空
//		return -1, errors.New("栈空")
//	}
//	return p.arr[p.stackSize-1], nil //栈顶元素
//}
//
////push栈元素
//func (p *Stack) push(t int) {
//	p.arr = append(p.arr, t)
//	p.stackSize = p.stackSize + 1
//}
//
////pop栈元素
//func (p *Stack) pop() (int, error) {
//	if p.stackSize > 0 { //栈不为空时
//		p.stackSize--
//		element := p.arr[p.stackSize]
//		p.arr = p.arr[:p.stackSize]
//		return element, nil
//	}
//	return -1, errors.New("栈空")
//}
//
//// 单增栈（出栈顺序单增）（失败）
//func maxArea__(height []int) int {
//	ret := 0
//	length := len(height)
//
//	// 初始化栈
//	stack := newStack()
//	stack.push(0)
//
//	// 元素下标依次入栈
//	// 当前元素大于栈顶元素时出栈
//	for i := 1; i < length; i++ {
//		if top, err := stack.peek(); err == nil && height[i] <= height[top] {
//			stack.push(i)
//		} else {
//			for {
//				val, _ := stack.pop()
//				ret = maxInt(ret, (i-val)*(minInt(height[i], height[val])))
//				if top, err := stack.peek(); err != nil || height[top] >= height[i] {
//					stack.push(i)
//					break
//				}
//			}
//		}
//	}
//
//	// 栈中元素处理
//	top, _ := stack.pop()
//	for {
//		if val, err := stack.pop(); err == nil {
//			ret = maxInt(ret, (top-val)*minInt(height[top], height[val]))
//		} else {
//			break
//		}
//	}
//	return ret
//}

//func maxInt(x int, y int) int {
//	if x < y {
//		return x
//	} else {
//		return y
//	}
//
//}
