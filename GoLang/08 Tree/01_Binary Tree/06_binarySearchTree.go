package main

import (
	"fmt"
)

type node struct {
	data   int
	left  *node
	right *node
}

// Create a node
func newNode(item int) *node {
	return &node{
		data: item, 
		left: nil, 
		right: nil,
	}
}

// Inorder Traversal
func inorder(root *node) {
	if root != nil {
		// Traverse left
		inorder(root.left)

		// Traverse root
		fmt.Printf("%d -> ", root.data)

		// Traverse right
		inorder(root.right)
	}
}

func search(root *node, data int) int {
	if root == nil {
		return -1
	} else if data == root.data {
		return root.data
	} else if data < root.data {
		return search(root.left, data)
	} else {
		return search(root.right, data)
	}
}

// Insert a node
func insert(root *node, data int) *node {
	// Return a new node if the tree is empty
	if root == nil {
		return newNode(data)
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
func findInorderSuccessor(node *node) *node {
	current := node

	// Find the leftmost leaf
	for current != nil && current.left != nil {
		current = current.left
	}

	return current
}

// Deleting a node
func deleteNode(root *node, data int) *node {
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

		// Place the inorder successor in position of the node to be deleted
		root.data = temp.data

		// Delete the inorder successor
		root.right = deleteNode(root.right, temp.data)
	}
	return root
}

func main() {
	var root *node
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

	fmt.Print("Inorder traversal: ")
	inorder(root)

	fmt.Println("\nAfter deleting 10")
	root = deleteNode(root, 1)
	fmt.Print("Inorder traversal: ")
	inorder(root)
}