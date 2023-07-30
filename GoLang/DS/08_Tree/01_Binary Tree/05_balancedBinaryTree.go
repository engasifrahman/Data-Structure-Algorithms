package main

import "fmt"

type Node struct {
	data int
	left *Node
	right *Node
}

// Height represents the height of a node.
type Height struct {
	height int
}

func CreateNode(data int) *Node {
	return &Node {
		data: data,
		left: nil,
		right: nil,
	}
}

func InsertLeft(rootNode *Node, leftNode *Node) bool {
	if rootNode.left != nil {
		fmt.Println("This node already assigned");
		return false
	}

	rootNode.left = leftNode

	return true
}

func InsertRight(rootNode *Node, rightNode *Node) bool {
	if rootNode.right != nil {
		fmt.Println("This node already assigned");
		return false
	}

	rootNode.right = rightNode
	
	return true
}

func DisplayPreorder(tree *Node) {
	fmt.Print("->", tree.data)

	if tree.left != nil {
		DisplayPreorder(tree.left)
	}

	if tree.right != nil {
		DisplayPreorder(tree.right)
	}
}

func DisplayInorder(tree *Node) {
	if tree.left != nil {
		DisplayInorder(tree.left)
	}

	fmt.Print("->", tree.data)

	if tree.right != nil {
		DisplayInorder(tree.right)
	}
}

func DisplayPostorder(tree *Node) {
	if tree.left != nil {
		DisplayPostorder(tree.left)
	}
	
	if tree.right != nil {
		DisplayPostorder(tree.right)
	}

	fmt.Print("->", tree.data)
}

// Check if the binary tree is height balanced.
func checkHeightBalance(root *Node, height *Height) bool {
	if root == nil {
		height.height = 0
		return true
	}

	var leftHeight, rightHeight Height

	l := checkHeightBalance(root.left, &leftHeight)
	r := checkHeightBalance(root.right, &rightHeight)

	fmt.Println("\nl: ", l)
	fmt.Println("r: ", r)

	leftHeightValue, rightHeightValue := leftHeight.height, rightHeight.height

	fmt.Println("leftHeightValue: ", leftHeightValue)
	fmt.Println("rightHeightValue: ", rightHeightValue)

	height.height = max(leftHeightValue, rightHeightValue) + 1

	fmt.Println("height.height: ", height.height)

	if abs(leftHeightValue - rightHeightValue) >= 2 {
		return false
	}

	return l && r
}


// Helper function to calculate the maximum of two integers.
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Helper function to calculate the absolute value of an integer.
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	root := CreateNode(1)
	node2 := CreateNode(2)
	node3 := CreateNode(3)
	node4 := CreateNode(4)
	node5 := CreateNode(5)
	node6 := CreateNode(6)

	InsertLeft(root, node2)
	InsertRight(root, node3)

	InsertLeft(node2, node4)
	InsertRight(node2, node5)

	InsertLeft(node5, node6)

	fmt.Print("DisplayPreorder")
	DisplayPreorder(root)

	fmt.Print("\nDisplayInorder")
	DisplayInorder(root)

	fmt.Print("\nDisplayPostorder")
	DisplayPostorder(root)

	var height Height
	if checkHeightBalance(root, &height) {
		fmt.Println("\nThe tree is balanced")
	} else {
		fmt.Println("\nThe tree is not balanced")
	}
}