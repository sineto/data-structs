package fifo

import "fmt"

type Item struct {
	Value int
}

type Queue struct {
	Items []*Item
	tail  int
}

func NewQueue() *Queue {
	return &Queue{}
}

func (item *Item) String() string {
	return fmt.Sprint(item.Value)
}

func (q *Queue) Push(item *Item) {
	q.Items = append(q.Items[:q.tail], item)
	q.tail = len(q.Items)
}

func (q *Queue) Shift() *Item {
	if q.tail == 0 {
		return nil
	}

	q.Items = append(q.Items[:0], q.Items[0+1:]...)
	q.tail = len(q.Items)

	if (q.tail - 1) < 0 {
		return nil
	}

	return q.Items[q.tail-1]
}
