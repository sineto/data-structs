package main

import (
	"fmt"

	"github.com/sineto/data-structs/binary_search"
	"github.com/sineto/data-structs/bubble"
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

	arr := []int{0, 1, 2, 3, 4, 5, 10, 1020, 1234}

	// Linear Search
	fmt.Println(linear_search.Search(arr, 10))
	fmt.Println(linear_search.Search(arr, 1))
	fmt.Println(linear_search.Search(arr, 0))
	fmt.Println(linear_search.Search(arr, 20))

	// Binary Search
	fmt.Println(binary_search.BinarySearch(arr, 10))
	fmt.Println(binary_search.BinarySearch(arr, 1234))
	fmt.Println(binary_search.BinarySearch(arr, 2))
	fmt.Println(binary_search.BinarySearch(arr, 20))

	arr = []int{8, 10, 9, 4, 7, 6, 0, 1}
	// Bubble Sort
	fmt.Println(arr)
	fmt.Println(bubble.BubbleSort(arr))
}
