package bubble

// Sort 排序
func Sort(a []int) []int {
	size := len(a)
	for i := 0; i < size-1; i++ {
		for j := 0; j < size-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}

	return a
}
