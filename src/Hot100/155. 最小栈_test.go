package Hot100

import "testing"

func TestMinStack(t *testing.T) {
	minStack := MinStackConstructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	t.Log(minStack.GetMin())
	minStack.Pop()
	t.Log(minStack.Top())
	t.Log(minStack.GetMin())
}
