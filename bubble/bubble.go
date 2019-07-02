package bubble

func SortAsc(data []int) {
	for i := 0; i < len(data)-1; i++ {
		flag := true
		for j := i + 1; j < len(data); j++ {
			if data[i] > data[j] {
				data[i], data[j] = data[j], data[i]
				flag = false
			}
		}
		if flag {
			break
		}
	}
}

func SortDesc(data []int) {
	for i := 0; i < len(data)-1; i++ {
		flag := true
		for j := i + 1; j < len(data); j++ {
			if data[i] < data[j] {
				data[i], data[j] = data[j], data[i]
				flag = false
			}
		}

		if flag {
			break
		}
	}
}
