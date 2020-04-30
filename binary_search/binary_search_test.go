package binary_search

import "testing"

func TestBinarySearch(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 10, 25, 46, 1024}

	assert := func(t *testing.T, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	}

	t.Run("return a right index value", func(t *testing.T) {
		got := BinarySearch(arr, 25)
		want := 7
		assert(t, got, want)
	})

	t.Run("return not founded value", func(t *testing.T) {
		got := BinarySearch(arr, 20)
		want := -1
		assert(t, got, want)
	})
}
