package stack

import "fmt"

type Stack struct {
	array []int
	top   int
}

func (stack *Stack) Push(val int) {
	if stack.array == nil{
		stack.array = []int{}
	}
	stack.array = append(stack.array, val)
	stack.top++
}

func (stack *Stack) Pop() {
	if stack.top == 0 {
		return
	}
	stack.array = stack.array[:stack.top-1]
	stack.top--
}

func (stack *Stack) Print() {
	fmt.Println(stack.array)
}

func (stack *Stack) Top() int{
	return stack.array[stack.top]
}

func (stack *Stack) Size() int{
	return stack.top
}