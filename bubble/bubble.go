package bubble

func BubbleSort(arr []int) []int {
	ordered := arr

	for i := 0; i < len(ordered); i++ {
		for j := 0; j < len(ordered)-1; j++ {
			if ordered[j] > ordered[j+1] {
				shadow := ordered[j]
				ordered[j] = ordered[j+1]
				ordered[j+1] = shadow
			}
		}
	}

	return ordered
}
