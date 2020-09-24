package tree

func InorderRecursive(node *TreeNode) (data []int) {
	if node == nil {
		return
	}
	data = append(data, InorderRecursive(node.Left)...)
	data = append(data, node.Val)
	data = append(data, InorderRecursive(node.Right)...)
	return
}

func InorderStack(node *TreeNode) (data []int) {
	stack := NewTreeNodeStack()
	for node != nil || stack.Len() > 0 {
		if node != nil {
			stack.Push(node)
			node = node.Left
		} else {
			back := stack.PopBack()
			data = append(data, back.Val)
			node = back.Right
		}
	}
	return
}

func InorderMorris(node *TreeNode) (data []int) {
	for node != nil {
		if node.Left != nil {
			predecessor := node.Left
			for predecessor.Right != nil && predecessor.Right != node {
				predecessor = predecessor.Right
			}
			if predecessor.Right == nil {
				predecessor.Right = node
				node = node.Left
			} else {
				data = append(data, node.Val)
				node = node.Right
				predecessor.Right = nil
			}
		} else {
			data = append(data, node.Val)
			node = node.Right
		}
	}
	return
}
