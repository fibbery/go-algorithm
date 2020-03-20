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
