package merge

// Sort 排序
func Sort(s []int) []int {
	size := len(s)
	if size <= 1 {
		return s
	}

	num := size / 2
	left := Sort(s[:num])
	right := Sort(s[num:])

	return merge(left, right)
}

// 合并
func merge(left, right []int) (result []int) {
	l, r := 0, 0
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}

	result = append(result, left[l:]...)
	result = append(result, right[r:]...)
	return
}
