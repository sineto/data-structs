package linear_search

func Search(arr []int, num int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == num {
			return i
		}
	}

	return -1
}
