package bubble

import (
	"testing"
)

func TestSortAsc(t *testing.T) {
	data := []int{1, 3, 2, 4}
	SortAsc(data)
	t.Log(data)
}

func TestSortDesc(t *testing.T) {
	data := []int{1, 3, 2, 4}
	SortDesc(data)
	t.Log(data)
}
