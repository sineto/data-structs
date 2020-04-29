package main

import (
	"fmt"

	"github.com/sineto/data-structs/fifo"
	"github.com/sineto/data-structs/lifo"
)

func main() {
	stack := lifo.NewStack()
	stack.Push(&lifo.Item{Value: 1})
	stack.Push(&lifo.Item{Value: 2})
	stack.Push(&lifo.Item{Value: 3})

	fmt.Println(stack.Items)
	for i := 0; i < 3; i++ {
		stack.Pop()
		fmt.Println(stack.Items)
	}

	queue := fifo.NewQueue()
	queue.Push(&fifo.Item{Value: 1})
	queue.Push(&fifo.Item{Value: 2})
	queue.Push(&fifo.Item{Value: 3})

	fmt.Println(queue.Items)
	for i := 0; i < 3; i++ {
		queue.Shift()
		fmt.Println(queue.Items)
	}
}
