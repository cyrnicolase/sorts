package selection

// Sort 排序
func Sort(a []int) []int {
	size := len(a)

	// 只需要排 size-1 次
	// 最后以此没有比较了，默认就是那个位置了
	for i := 0; i < size-1; i++ {
		k := i
		// 排序位置是待排序序列的第二个元素；j := i + 1
		for j := i + 1; j < size; j++ {
			if a[k] > a[j] {
				k = j
			}
		}
		a[k], a[i] = a[i], a[k]
	}

	return a
}
