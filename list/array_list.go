package list

type ArrayList struct {
	head *Node
	size int
}

func NewArrayList() *ArrayList {
	return &ArrayList{}
}

type Node struct {
	Val  interface{}
	Next *Node
}

func (l *ArrayList) Push(val interface{}) {
	next := &Node{Val: val}
	if l.head == nil {
		l.head = next
		return
	}
	node := l.head
	for node.Next != nil {
		node = node.Next
	}
	node.Next = next
}

func (l *ArrayList) AddFirst(val interface{}) {
	l.size++
	node := &Node{Val: val}
	if l.head == nil {
		l.head = node
		return
	}
	node.Next = l.head
	l.head = node
}

func (l *ArrayList) Size() int {
	return l.size
}
