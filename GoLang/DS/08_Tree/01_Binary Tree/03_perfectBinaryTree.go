package main

import "fmt"

type Node struct {
	data int
	left *Node
	right *Node
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

// findTreeHeight finds the height of the binary tree.
func findTreeHeight(root *Node) int {
	if root == nil {
		return 0
	}

	leftHeight := findTreeHeight(root.left)

	rightHeight := findTreeHeight(root.right)

	if leftHeight > rightHeight {
		return leftHeight + 1
	}

	return rightHeight + 1
}

// isLeaf checks if a node is a leaf (i.e., has no children).
func isLeaf(node *Node) bool {
	return node != nil && node.left == nil && node.right == nil
}

// isPerfectBinaryTree checks if the given binary tree is perfect.
func isPerfectBinaryTree(root *Node) bool {
	if root == nil {
		return true
	}

	height := findTreeHeight(root)
	return checkPerfect(root, height, 0)
}

// checkPerfect recursively checks if the binary tree is perfect.
func checkPerfect(node *Node, height, level int) bool {
	if node == nil {
		return true
	}

	// Check if the node is a leaf.
	if isLeaf(node) {
		return height == level+1
	}

	// Check if the node is an internal node and has two children.
	if node.left == nil || node.right == nil {
		return false
	}

	// Recursively check the left and right subtrees.
	return checkPerfect(node.left, height, level+1) && checkPerfect(node.right, height, level+1)
}


func main() {
	root := CreateNode(1)
	node2 := CreateNode(2)
	node3 := CreateNode(3)
	node4 := CreateNode(4)
	node5 := CreateNode(5)
	node6 := CreateNode(6)
	node7 := CreateNode(7)

	InsertLeft(root, node2)
	InsertRight(root, node3)

	InsertLeft(node2, node4)
	InsertRight(node2, node5)

	InsertLeft(node2, node6)
	InsertRight(node2, node7)

	fmt.Print("DisplayPreorder")
	DisplayPreorder(root)

	fmt.Print("\nDisplayInorder")
	DisplayInorder(root)

	fmt.Print("\nDisplayPostorder")
	DisplayPostorder(root)

	fmt.Println("IsFullBinaryTree:", isPerfectBinaryTree(root))
}