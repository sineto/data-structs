package fifo

import (
	"reflect"
	"testing"
)

func TestFifo(t *testing.T) {
	got := NewQueue()
	got.Push(&Item{1})
	got.Push(&Item{2})

	assertEqual := func(t *testing.T, got, want []*Item) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	}

	t.Run("push an item into queue", func(t *testing.T) {
		want := []*Item{{1}, {2}}
		assertEqual(t, got.Items, want)
	})

	t.Run("shift an item from a head of queue", func(t *testing.T) {
		got.Shift()
		want := []*Item{{2}}
		assertEqual(t, got.Items, want)
	})
}
