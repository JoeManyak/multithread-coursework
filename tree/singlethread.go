package tree

func (n *Node) SingleThreadSearch() *Node {
	return n.singleThreadSearchUtil()
}

func (n *Node) singleThreadSearchUtil() *Node {
	if n.IsNeeded {
		return n
	}
	n.Visited = true

	if n.Left != nil {
		if !n.Left.Visited {
			node := n.Left.singleThreadSearchUtil()
			if node != nil {
				return node
			}
		}
	}
	if n.Right != nil {
		if !n.Right.Visited {
			node := n.Right.singleThreadSearchUtil()
			if node != nil {
				return node
			}
		}
	}

	return nil
}
