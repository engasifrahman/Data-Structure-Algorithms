package main

import (
	"fmt"
)

type Node struct {
	data   int
	left  *Node
	right *Node
}

// Create a node
func createNode(item int) *Node {
	return &Node{
		data: item, 
		left: nil, 
		right: nil,
	}
}

// InorderTraversal Traversal
func inorderTraversal(root *Node) {
	if root != nil {
		// Traverse left
		inorderTraversal(root.left)

		// Traverse root
		fmt.Printf("%d -> ", root.data)

		// Traverse right
		inorderTraversal(root.right)
	}
}

func search(root *Node, data int) bool {
	if root == nil {
		return false
	} else if data == root.data {
		return true
	} else if data < root.data {
		return search(root.left, data)
	} else {
		return search(root.right, data)
	}
}

// Insert a node
func insert(root *Node, data int) *Node {
	// Return a new node if the tree is empty
	if root == nil {
		return createNode(data)
	}

	// Traverse to the right place and insert the node
	if data < root.data {
		root.left = insert(root.left, data)
	} else {
		root.right = insert(root.right, data)
	}

	return root
}

// Find the inorder successor
func findInorderSuccessor(node *Node) *Node {
	current := node

	// Find the leftmost leaf
	for current != nil && current.left != nil {
		current = current.left
	}

	return current
}

// Deleting a node
func deleteNode(root *Node, data int) *Node {
	// Return if the tree is empty
	if root == nil {
		return root
	}

	// Find the node to be deleted
	if data < root.data {
		root.left = deleteNode(root.left, data)
	} else if data > root.data {
		root.right = deleteNode(root.right, data)
	} else {
		// If the node is with only one child or no child
		if root.left == nil {
			temp := root.right
			root = nil
			return temp
		} else if root.right == nil {
			temp := root.left
			root = nil
			return temp
		}

		// If the node has two children
		temp := findInorderSuccessor(root.right)

		// Place the inorderTraversal successor in position of the node to be deleted
		root.data = temp.data

		// Delete the inorderTraversal successor
		root.right = deleteNode(root.right, temp.data)
	}
	return root
}

func main() {
	var root *Node
	root = insert(root, 8)
	root = insert(root, 3)
	root = insert(root, 1)
	root = insert(root, 6)
	root = insert(root, 7)
	root = insert(root, 10)
	root = insert(root, 14)
	root = insert(root, 4)

	sr := search(root, 10)
	fmt.Println("Search Result: ", sr)

	fmt.Print("Inorder Traversal: ")
	inorderTraversal(root)

	fmt.Println("\nAfter deleting")
	root = deleteNode(root, 9)

	fmt.Print("Inorder Traversal: ")
	inorderTraversal(root)
}