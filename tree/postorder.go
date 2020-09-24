package tree

func PostorderRecursive(node *TreeNode) (data []int) {
	if node == nil {
		return
	}
	data = append(data, PostorderRecursive(node.Left)...)
	data = append(data, PostorderRecursive(node.Right)...)
	data = append(data, node.Val)
	return
}

func PostorderStack(node *TreeNode) (data []int) {
	stack := NewTreeNodeStack()
	var previous *TreeNode
	for node != nil || stack.Len() > 0 {
		for node != nil {
			stack.Push(node)
			node = node.Left
		}
		node = stack.Back()
		if (node.Left == nil && node.Right == nil) ||
			(node.Right == nil && node.Left == previous) ||
			(node.Right == previous) {
			data = append(data, node.Val)
			previous = node
			stack.PopBack()
			node = nil
		} else {
			node = node.Right
		}

	}
	return
}
