package sort

import "testing"

func Test_insertSortAsc(t *testing.T) {
	data := []int{1, 3, 4, 2, 5, 6,}
	insertSortAsc(data)
	t.Log(data)
}

func Test_insertSortDesc(t *testing.T) {
	data := []int{1, 3, 4, 2, 5, 6,}
	insertSortDesc(data)
	t.Log(data)
}
