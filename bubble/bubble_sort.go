package bubble

// Sort 排序
func Sort(a []int) []int {
	size := len(a)
	for i := 0; i < size; i++ {
		// 比较响铃元素的最大上限是 size-i-1;
		// 因为每一轮排序都将待排序的最大值交换到最后
		// 所以后面的轮次的最大元素都不会比上一次的更大
		for j := 0; j < size-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}

	return a
}
