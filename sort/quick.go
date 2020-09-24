package sort

import "container/list"

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


// 第二种解法
func sortArray(nums []int) []int {
	quickSort2(nums, 0, len(nums) - 1)
	return nums
}

func quickSort2(nums []int, left, right int) {
	if left < right {
		pos := partition(nums, left, right)
		quickSort2(nums, left, pos - 1)
		quickSort2(nums, pos + 1, right)
	}
}

func partition(nums []int, left, right int)int {
	pivot := left
	//从基准值后一个序号开始处理
	index := pivot + 1
	for j := index ; j <= right; j ++ {
		if nums[j] <= nums[pivot] {
			nums[j], nums[index] = nums[index], nums[j]
			index++
		}
	}
	nums[pivot] , nums[index - 1] = nums[index - 1] , nums[pivot]
	return index - 1

	list.New()
}


