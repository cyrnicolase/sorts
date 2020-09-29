package insertion

import (
	"reflect"
	"testing"
)

// TestSort 测试插入排序
func TestSort(t *testing.T) {
	a := []int{3, 1, 5, 2, 7, 1, 9, 4, 8, 2}
	ea := []int{1, 1, 2, 2, 3, 4, 5, 7, 8, 9}

	ra := Sort(a)
	if reflect.DeepEqual(ea, ra) {
		t.Log("插入排序算法通过！")
		return
	}

	t.Errorf("插入排序算法不正确\n原始序列：%v\n期望输出：%v\n实际输出：%v\n", a, ea, ra)
}
