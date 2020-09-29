package list

import "testing"

func TestNewArrayList(t *testing.T) {
	list := NewArrayList()
	list.Push(1)
	list.Push(2)
	if list.Size() != 2{
		t.Errorf("list size must equal 2")
	}
	list.AddFirst(3)
	if list.Head().Val.(int) != 3{
		t.Errorf("list head must euqal 3")
	}
}
