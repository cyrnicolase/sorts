package shell

// Sort 排序
func Sort(s []int) []int {
	size := len(s)
	for gap := size >> 1; gap > 0; gap >>= 1 {
		for i := gap; i < size; i++ {
			temp := s[i]

			var j int
			for j = i - gap; j >= 0 && s[j] > temp; j -= gap {
				s[j+gap] = s[j]
			}
			s[j+gap] = temp
		}
	}

	return s
}
