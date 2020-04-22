package sort

func HeapSort(array []int) []int {
	length := len(array)
	beginIndex := length/2 - 1
	for i := beginIndex; i >= 0; i-- {
		//创造出一个大顶堆：即节点永远比该节点的两个子节点大
		heapify(array, i, length)
	}
	//堆排序
	for i := length - 1; i > 0; i-- {
		//将大顶堆的最大的元素（即堆顶元素）和最后一个元素进行置换。这样保证最大数在数组最后一个槽位
		array[0], array[i] = array[i], array[0]
		heapify(array, 0, i)
	}
	return array
}

func heapify(array []int, index int, length int) {
	// 左节点索引
	left := 2*index + 1
	// 右节点索引
	right := left + 1
	// 子节点最大值的索引 默认为左节点
	largest := left

	//获取子节点中最大值的索引
	if left >= length {
		return
	}
	if right < length && array[right] > array[left] {
		largest = right
	}

	//如果子节点中最大值超过该节点，那么进行替换
	//之后，下沉到和该节点交换值的节点进行再一次的堆化，保证堆正确性
	if array[largest] > array[index] {
		array[largest], array[index] = array[index], array[largest]
		heapify(array, largest, length)
	}
}
