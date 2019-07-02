package selection

import "testing"

func TestSortAsc(t *testing.T) {
	data := []int{1,3,4,2,6,5,}
	SortAsc(data)
	t.Log(data)
}

func TestSortDesc(t *testing.T) {
	data := []int{1,3,4,2,6,5,}
	SortDesc(data)
	t.Log(data)
}
