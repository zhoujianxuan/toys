package design

import (
	"fmt"
	"testing"
)

func TestMinStack(t *testing.T) {
	stack := Constructor()
	stack.Push(-2)
	stack.Push(0)
	stack.Push(-1)
	fmt.Println(stack.GetMin())
	fmt.Println(stack.Top())
	stack.Pop()
	fmt.Println(stack.GetMin())
}
