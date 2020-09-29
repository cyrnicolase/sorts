package insertion

// Sort 排序
func Sort(s []int) []int {
	size := len(s)
	if size <= 1 {
		return s
	}

	for i := 1; i < size; i++ {
		for j := i; j > 0; j-- {
			if s[j] < s[j-1] {
				s[j-1], s[j] = s[j], s[j-1]
			}
		}
	}

	return s
}
