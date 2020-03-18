package sort

import "testing"

func Test_selectSortAsc(t *testing.T) {
	data := []int{1,3,4,2,6,5,}
	selectSortAsc(data)
	t.Log(data)
}

func Test_selectSortDesc(t *testing.T) {
	data := []int{1,3,4,2,6,5,}
	selectSortDesc(data)
	t.Log(data)
}
