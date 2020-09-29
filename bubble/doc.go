package bubble

// 冒泡排序算法是把较小的元素往前调或把较大的元素往后调。这种方法主要是通过对相临两个元素
// 进行大小的比较，根据比较结果和算法规则对该二元素的位置进行交换，这样逐个依次进行比较和
// 交换，就能达到排序目的。冒泡排序的基本思想是，首先将第一个和第二个记录的关键字比较大小，
// 如果是逆序的，就将这两个记录进行交换，再对第2个和第3个记录的关键字进行比较，依次类推，
// 重复进行上述计算，直至完成第(n-1)个和第n个记录的关键字之间比较，此后，再按照上述过程进
// 行第2次、第3次排序，直至整个序列有序为止。排序过程中要特别注意的是，当响铃两个元素大小一
// 致时，这一步骤就不需要交换位置，因此页说明冒泡排序是一种严格的稳定排序算法，它不改变序列
// 中相同元素之间的相对位置关系。

// 冒泡排序算法的时间复杂度是： O(n^2)