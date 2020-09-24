package heap

type Comparable struct {
	Val int
}

func (c Comparable) Compare(b Comparable) int {
	if c.Val > b.Val {
		return 1
	} else if c.Val < b.Val {
		return -1
	} else {
		return 0
	}
}

type BinaryHeap struct {
	data []Comparable
	size int
}

func NewBinaryHeap() *BinaryHeap {
	return &BinaryHeap{
		data: make([]Comparable, 0),
		size: 0,
	}
}

func (b *BinaryHeap) Insert(value Comparable) {
	b.data = append(b.data, value)
	b.size++
	b.swim(b.size - 1)
}

func (b *BinaryHeap) Pop() Comparable {
	el := b.data[0]
	b.data[b.size-1], b.data[0] = b.data[0], b.data[b.size-1]
	b.data = b.data[:b.size-1]
	b.size--
	b.sink(0)
	return el
}

func (b *BinaryHeap) swim(index int) {
	for index >= 0 && b.data[parent(index)].Compare(b.data[index]) < 0 {
		b.data[parent(index)], b.data[index] = b.data[index], b.data[parent(index)]
		b.swim(parent(index))
	}
}

func (b *BinaryHeap) sink(index int) {
	if index < b.size {
		largest := left(index)
		if largest > b.size-1 {
			return
		}
		if right(index) < b.size-1 && b.data[right(index)].Compare(b.data[largest]) > 0 {
			largest = right(index)
		}

		if b.data[largest].Compare(b.data[index]) <= 0 {
			return
		}
		b.data[largest], b.data[index] = b.data[index], b.data[largest]
		b.sink(largest)
	}
}

//父节点索引
func parent(index int) int {
	if index == 0 {
		return 0
	}
	if isOdd(index) {
		return index / 2
	}
	return index/2 - 1
}

func left(index int) int {
	return 2*index + 1
}

func right(index int) int {
	return 2*index + 2
}

func isOdd(index int) bool {
	if index%2 == 0 {
		return false
	}
	return true
}
