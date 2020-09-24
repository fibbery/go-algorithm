package tree

func PreorderRecursive(node *TreeNode) (data []int) {
	if node == nil {
		return
	}
	data = append(data, node.Val)
	data = append(data, PreorderRecursive(node.Left)...)
	data = append(data, PreorderRecursive(node.Right)...)
	return data
}

func PreorderStack(node *TreeNode) (data []int) {
	stack := NewTreeNodeStack()
	for node != nil || stack.Len() > 0 {
		if node != nil {
			data = append(data, node.Val)
			stack.Push(node)
			node = node.Left
		} else {
			back := stack.PopBack()
			node = back.Right
		}
	}
	return data
}

func PreorderMorris(node *TreeNode) (data []int) {
	for node != nil {
		if node.Left != nil {
			predecessor := node.Left

			//寻找前驱结点
			for predecessor.Right != nil && predecessor.Right != node {
				predecessor = predecessor.Right
			}

			if predecessor.Right == nil {
				data = append(data, node.Val)
				predecessor.Right = node
				node = node.Left
			} else {
				node = node.Right
				predecessor.Right = nil
			}

		} else {
			data = append(data, node.Val)
			node = node.Right
		}
	}
	return data
}
