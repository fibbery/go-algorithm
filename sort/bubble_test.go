package sort

import (
	"testing"
)


func Test_bubbleSortAsc(t *testing.T) {
	data := []int{1, 3, 2, 4}
	bubbleSortAsc(data)
	t.Log(data)
}

func Test_bubbleSortDesc(t *testing.T) {
	data := []int{1, 3, 2, 4}
	bubbleSortDesc(data)
	t.Log(data)
}