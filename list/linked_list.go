package list

type LinkList struct {
	head *LinkNode
	size int
}

func NewLinkList() *LinkList {
	return &LinkList{}
}

type LinkNode struct {
	Val  interface{}
	Next *LinkNode
}

func (l *LinkList) Push(val interface{}) {
	next := &LinkNode{Val: val}
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

func (l *LinkList) AddFirst(val interface{}) {
	l.size++
	node := &LinkNode{Val: val}
	if l.head == nil {
		l.head = node
		return
	}
	node.Next = l.head
	l.head = node
}

func (l *LinkList) Size() int {
	return l.size
}
