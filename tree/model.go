package tree

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

type TreeNodeStack struct {
	data []*TreeNode
}

func NewTreeNodeStack() *TreeNodeStack {
	return &TreeNodeStack{data: make([]*TreeNode, 0)}
}

func (t *TreeNodeStack) Push(node *TreeNode) {
	t.data = append(t.data, node)
}

func (t *TreeNodeStack) PopBack() *TreeNode {
	last := t.data[len(t.data)-1]
	t.data = t.data[:len(t.data)-1]
	return last
}

func (t *TreeNodeStack) Back() *TreeNode {
	return t.data[len(t.data)-1]
}

func (t *TreeNodeStack) Len() int {
	return len(t.data)
}

func CreateTree(data []int, index, length int) *TreeNode {
	if index >= length {
		return nil
	}

	node := &TreeNode{Val: data[index]}
	node.Left = CreateTree(data, 2*index+1, length)
	node.Right = CreateTree(data, 2*index+2, length)
	return node
}
