package main

import (
	"fmt"

	"github.com/sineto/data-structs/fifo"
	"github.com/sineto/data-structs/lifo"
	"github.com/sineto/data-structs/linear_search"
)

func main() {
	// LIFO structure
	stack := lifo.NewStack()
	stack.Push(&lifo.Item{Value: 1})
	stack.Push(&lifo.Item{Value: 2})
	stack.Push(&lifo.Item{Value: 3})

	fmt.Println(stack.Items)
	for i := 0; i < 3; i++ {
		stack.Pop()
		fmt.Println(stack.Items)
	}

	// FIFO structure
	queue := fifo.NewQueue()
	queue.Push(&fifo.Item{Value: 1})
	queue.Push(&fifo.Item{Value: 2})
	queue.Push(&fifo.Item{Value: 3})

	fmt.Println(queue.Items)
	for i := 0; i < 3; i++ {
		queue.Shift()
		fmt.Println(queue.Items)
	}

	// Linear Search
	arr := []int{1, 2, 3, 4, 5, 10, 1020, 1234, 0}
	fmt.Println(linear_search.Search(arr, 10))
	fmt.Println(linear_search.Search(arr, 1))
	fmt.Println(linear_search.Search(arr, 0))
}
