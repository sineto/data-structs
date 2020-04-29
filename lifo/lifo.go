package lifo

import "fmt"

type Item struct {
	Value int
}

type Stack struct {
	Items []*Item
	top   int
}

func NewStack() *Stack {
	return &Stack{}
}

func (item *Item) String() string {
	return fmt.Sprint(item.Value)
}

func (s *Stack) Push(item *Item) {
	s.Items = append(s.Items[:s.top], item)
	s.top++
}

func (s *Stack) Pop() *Item {
	if s.top == 0 {
		return nil
	}

	s.Items = s.Items[:len(s.Items)-1]
	s.top--

	if (s.top - 1) < 0 {
		return nil
	}

	return s.Items[s.top-1]
}
