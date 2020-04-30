package bubble

import (
	"reflect"
	"testing"
)

func TestBubble(t *testing.T) {
	arr := []int{8, 10, 9, 4, 7, 6, 0, 1}
	got := BubbleSort(arr)
	want := []int{0, 1, 4, 6, 7, 8, 9, 10}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
