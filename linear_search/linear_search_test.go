package linear_search

import "testing"

func TestLinearSearch(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6}
	assert := func(t *testing.T, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	}

	t.Run("find an item", func(t *testing.T) {
		got := Search(arr, 4)
		want := 3
		assert(t, got, want)
	})

	t.Run("can't find an item", func(t *testing.T) {
		got := Search(arr, 10)
		want := -1
		assert(t, got, want)
	})
}
