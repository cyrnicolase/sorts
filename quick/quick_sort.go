package quick

// Sort 排序
func Sort(s []int) {
	size := len(s)
	if size <= 1 {
		return
	}

	quickSort(s, 0, size-1)
}

// 快速排序
func quickSort(s []int, left, right int) {
	p, i, j := left, left, right
	value := s[left]
	for i <= j {
		for j >= p && s[j] >= value {
			j--
		}
		if j >= p {
			s[p] = s[j]
			p = j
		}

		for i <= p && s[i] <= value {
			i++
		}
		if i <= p {
			s[p] = s[i]
			p = i
		}
	}

	s[p] = value
	if p-left > 1 {
		quickSort(s, left, p-1)
	}
	if right-p > 1 {
		quickSort(s, p+1, right)
	}
}
