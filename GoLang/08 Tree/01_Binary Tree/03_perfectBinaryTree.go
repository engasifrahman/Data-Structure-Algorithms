package main

import "fmt"

type Node struct {
	data int
	left *Node
	right *Node
}

func createNode(data int) *Node {
	return &Node {
		data: data,
		left: nil,
		right: nil,
	}
}

func insertLeft(rootNode *Node, leftNode *Node) bool {
	if rootNode.left != nil {
		fmt.Println("This node already assigned");
		return false
	}

	rootNode.left = leftNode

	return true
}

func insertRight(rootNode *Node, rightNode *Node) bool {
	if rootNode.right != nil {
		fmt.Println("This node already assigned");
		return false
	}

	rootNode.right = rightNode
	
	return true
}

func preorderTraversal(tree *Node) {
	fmt.Print("->", tree.data)

	if tree.left != nil {
		preorderTraversal(tree.left)
	}

	if tree.right != nil {
		preorderTraversal(tree.right)
	}
}

func inorderTraversal(tree *Node) {
	if tree.left != nil {
		inorderTraversal(tree.left)
	}

	fmt.Print("->", tree.data)

	if tree.right != nil {
		inorderTraversal(tree.right)
	}
}

func postorderTraversal(tree *Node) {
	if tree.left != nil {
		postorderTraversal(tree.left)
	}
	
	if tree.right != nil {
		postorderTraversal(tree.right)
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

	return max(leftHeight, rightHeight) + 1
}

// Find maximum number
func max(a int, b int) int {
	if a > b {
		return a
	}

	return b
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
	root := createNode(1)
	node2 := createNode(2)
	node3 := createNode(3)
	node4 := createNode(4)
	node5 := createNode(5)
	node6 := createNode(6)
	node7 := createNode(7)

	insertLeft(root, node2)
	insertRight(root, node3)

	insertLeft(node2, node4)
	insertRight(node2, node5)

	insertLeft(node3, node6)
	insertRight(node3, node7)

	fmt.Print("Preorder Traversal")
	preorderTraversal(root)

	fmt.Print("\nInorder Traversal")
	inorderTraversal(root)

	fmt.Print("\nPostorder Traversal")
	postorderTraversal(root)

	fmt.Println("\nIs Perfect Binary Tree:", isPerfectBinaryTree(root))
}