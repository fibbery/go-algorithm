# go-algorithm
算法（go实现）

## 经典排序算法
### 1.冒泡排序
#### 1.1 算法步骤
* 比较相邻的元素。如果第一个比第二个大，就交换他们两个。
* 对每一对相邻元素作同样的工作，从开始第一对到结尾的最后一对。这步做完后，最后的元素会是最大的数。
* 针对所有的元素重复以上的步骤，除了最后一个。
* 持续每次对越来越少的元素重复上面的步骤，直到没有任何一对数字需要比较。
#### 1.2 参考代码
```go
package sort

func bubbleSortAsc(data []int) {
	for i := 0; i < len(data)-1; i++ {
		for j := i + 1; j < len(data); j++ {
			if data[i] > data[j] {
				data[i], data[j] = data[j], data[i]
			}
		}
	}
}

func bubbleSortDesc(data []int) {
	for i := 0; i < len(data)-1; i++ {
		for j := i + 1; j < len(data); j++ {
			if data[i] < data[j] {
				data[i], data[j] = data[j], data[i]
			}
		}
	}
}
```
### 2.选择排序
#### 2.1 算法步骤
* 首先在未排序序列中找到最小（大）元素，存放到排序序列的起始位置
* 再从剩余未排序元素中继续寻找最小（大）元素，然后放到已排序序列的末尾。
* 重复第二步，直到所有元素均排序完毕。
#### 2.2 参考代码
```go
package sort

func selectSortAsc(data []int) {
	for i := 0; i < len(data)-1; i++ {
		min := i
		for j := i + 1; j < len(data); j++ {
			if data[j] < data[min] {
				min = j
			}
		}
		if min != i {
			data[i], data[min] = data[min], data[i]
		}
	}
}

func selectSortDesc(data []int) {
	for i := 0; i < len(data)-1; i++ {
		max := i
		for j := i + 1; j < len(data); j++ {
			if data[j] > data[max] {
				max = j
			}
		}
		if max != i {
			data[i], data[max] = data[max], data[i]
		}
	}
}
```
### 3.插入排序
#### 3.1 算法步骤
* 将第一待排序序列第一个元素看做一个有序序列，把第二个元素到最后一个元素当成是未排序序列。
* 从头到尾依次扫描未排序序列，将扫描到的每个元素插入有序序列的适当位置。（如果待插入的元素与有序序列中的某个元素相等，则将待插入元素插入到相等元素的后面。）
#### 3.2 参考代码
```go
package sort

func insertSortAsc(data []int) {
	for i := 1; i < len(data); i++ {
		val := data[i] //需要插入的数
		key := i - 1
		//如果插入的数比左侧的小，则把左侧的数前移一位
		for ; key >= 0 && val < data[key]; key-- {
			data[key+1] = data[key]
		}
		// 将插入的数放在不小于左侧数的位置
		if key+1 != i {
			data[key+1] = val
		}
	}
}

func insertSortDesc(data []int) {
	for i := 1; i < len(data); i++ {
		val := data[i]
		key := i - 1
		for ;key >= 0 && val > data[key];key-- {
			data[key+1] = data[key]
		}

		if key != i-1 {
			data[key+1] = val
		}
	}
}
```
### 4.快速排序
#### 4.1 算法步骤
* 从数列中挑出一个元素，称为 “基准”（pivot）;
* 重新排序数列，所有元素比基准值小的摆放在基准前面，所有元素比基准值大的摆在基准的后面（相同的数可以到任一边）。在这个分区退出之后，该基准就处于数列的中间位置。这个称为分区（partition）操作；
* 递归地（recursive）把小于基准值元素的子数列和大于基准值元素的子数列排序；
#### 4.2 参考代码
```go
package sort

func quickSort(array []int) []int {
	if len(array) < 2 {
		return array
	}
	pivot := array[0]
	var left, middle, right []int
	for _, v := range array {
		if v > pivot {
			right = append(right, v)
		}else if v == pivot {
			middle = append(middle, v)
		}else {
			left = append(left, v)
		}
	}
	return append(quickSort(append(left, middle...)), quickSort(right)...)
}
```
### 5 堆排序
#### 5.1 算法步骤
1. 创建一个大顶堆
2. 把大顶堆的第一个元素同最后一个元素交换
3. 缩小堆的尺寸，并继续进行堆化，目的是为了继续把最大的元素调整到顶端
4. 从第二部开始循环，一直到堆的尺寸缩小到1为止
#### 5.2 参考代码
```go
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

```