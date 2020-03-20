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
