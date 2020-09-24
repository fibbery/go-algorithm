package heap

import (
	"sort"
	"testing"
)

func TestBinaryHeap(t *testing.T) {
	heap := NewBinaryHeap()
	data := []int{5, 9, 4, 2, 6, 9}
	for _, value := range data {
		heap.Insert(Comparable{Val: value})
	}
	//排序一下data
	sort.Ints(data)
	for i := len(data) - 1; i >= 0; i-- {
		pop := heap.Pop()
		if pop.Val != data[i] {
			t.Fatalf("pop value [%d] must equal array value [%d]", pop.Val, data[i])
		}
	}
}
