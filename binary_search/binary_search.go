package binary_search

func BinarySearch(arr []int, num int) int {
	left := 0
	right := len(arr) - 1

	if right < left {
		return -1
	}

	for left <= right {
		mid := (left + right) / 2

		if arr[mid] == num {
			return mid
		} else if arr[mid] < num {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}
