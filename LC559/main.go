package main

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	maxChildDepth := 0
	for _, child := range root.Children {
		if childDepth := maxDepth(child); childDepth > maxChildDepth {
			maxChildDepth = childDepth
		}
	}
	return maxChildDepth + 1
}
