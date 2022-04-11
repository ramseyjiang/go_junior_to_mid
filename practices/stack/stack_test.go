package stack

import (
	"fmt"
	"testing"
)

func TestMinStack(t *testing.T) {
	ms := New()
	fmt.Println("maxInt64 is ", ms.GetMin())

	ms.Push(20)
	ms.Push(39)
	fmt.Println("MinStack content after push 39 is: ", ms.stack)
	fmt.Println("MinStack after push 39, the min is: ", ms.GetMin())

	ms.Push(10)
	fmt.Println("MinStack content after push 10 is: ", ms.stack)
	fmt.Println("MinStack after push 10, the min is: ", ms.GetMin())

	fmt.Println(ms.GetLastElement())
	fmt.Println("MinStack content after pop is: ", ms.stack)
	ms.Pop()
	fmt.Println("MinStack content after pop is: ", ms.stack)

}
