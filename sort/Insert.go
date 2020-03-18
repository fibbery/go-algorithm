package sort

func insertSortAsc(data []int) {
	for i := 1; i < len(data); i++ {
		for j := i - 1; j >= 0; j-- {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
}

func insertSortDesc(data []int) {
	for i := 1; i < len(data); i++ {
		for j := i - 1; j >= 0; j-- {
			if data[j] < data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
}
