package lifo

import (
	"reflect"
	"testing"
)

func TestLifo(t *testing.T) {
	got := NewStack()
	got.Push(&Item{1})
	got.Push(&Item{2})

	assertEqual := func(t *testing.T, got, want []*Item) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	}

	t.Run("push an item into stack", func(t *testing.T) {
		want := []*Item{{1}, {2}}
		assertEqual(t, got.Items, want)

	})

	t.Run("pop an item from stack", func(t *testing.T) {
		got.Push(&Item{3})
		want := []*Item{{1}, {2}}
		got.Pop()
		assertEqual(t, got.Items, want)
	})
}
