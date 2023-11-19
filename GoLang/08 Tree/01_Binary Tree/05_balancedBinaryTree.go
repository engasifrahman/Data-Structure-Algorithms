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

// Check if the binary tree is height balanced.
func isBalancedBinaryTree(node *Node) (bool, int) {
	if node == nil {
		return true, 0
	}

	left, leftHeight := isBalancedBinaryTree(node.left)
	right, rightHeight := isBalancedBinaryTree(node.right)

	height := max(leftHeight, rightHeight) + 1

	if abs(leftHeight - rightHeight) > 1 {
		return false, height
	}

	return left && right, height
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
	root := createNode(1)
	node2 := createNode(2)
	node3 := createNode(3)
	node4 := createNode(4)
	node5 := createNode(5)
	node6 := createNode(6)

	insertLeft(root, node2)
	insertRight(root, node3)

	insertLeft(node2, node4)
	insertRight(node2, node5)

	insertLeft(node3, node5)
	insertLeft(node4, node6)

	fmt.Print("Preorder Traversal")
	preorderTraversal(root)

	fmt.Print("\nInorder Traversal")
	inorderTraversal(root)

	fmt.Print("\nPostorder Traversal")
	postorderTraversal(root)

	flag,_ := isBalancedBinaryTree(root)
	if flag {
		fmt.Println("\nThe tree is balanced")
	} else {
		fmt.Println("\nThe tree is not balanced")
	}
}