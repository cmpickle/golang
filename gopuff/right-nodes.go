package main

type BinaryTree struct {
	Value int
	Left  *BinaryTree
	Right *BinaryTree
}

type Node struct {
	tree   *BinaryTree
	depth  int
	xValue int
}

func treeRightNodes(root *BinaryTree) []int {
	queue := []*Node{{tree: root, depth: 0, xValue: 0}}
	var rightMostValue map[int]*Node

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if rightMostValue[curr.depth].xValue < curr.xValue {
			rightMostValue[curr.depth] = curr
		}
		if curr.tree.Left != nil {
			queue = append(queue, &Node{tree: curr.tree.Left, depth: curr.depth + 1, xValue: curr.xValue - 1})
		}
		if curr.tree.Right != nil {
			queue = append(queue, &Node{tree: curr.tree.Left, depth: curr.depth + 1, xValue: curr.xValue + 1})
		}
	}

	rightValues := []int{}
	for _, node := range rightMostValue {
		rightValues = append(rightValues, node.tree.Value)
	}

	return rightValues
}
