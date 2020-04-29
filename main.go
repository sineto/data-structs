package main

import (
	"fmt"

	"github.com/sineto/data-structs/lifo"
)

func main() {
	stack := lifo.NewStack()
	stack.Push(&lifo.Item{Value: 1})
	stack.Push(&lifo.Item{Value: 2})
	stack.Push(&lifo.Item{Value: 3})

	fmt.Println(stack.Items)
	stack.Pop()
	fmt.Println(stack.Items)
	stack.Pop()
	fmt.Println(stack.Items)
	stack.Pop()
	fmt.Println(stack.Items)
}
