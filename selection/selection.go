package selection

func SortAsc(data []int) {
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

func SortDesc(data []int) {
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
