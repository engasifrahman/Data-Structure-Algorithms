//? Full binary tree
// A full binary tree is a special type of binary tree in which every parent node/internal node has either two or no children.
// It is also known as a Proper Binary tree.

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

func insertLeft(parentNode *Node, leftNode *Node) bool {
	if parentNode.left != nil {
		fmt.Println("This node already assigned");
		return false
	}

	parentNode.left = leftNode

	return true
}

func insertRight(parentNode *Node, rightNode *Node) bool {
	if parentNode.right != nil {
		fmt.Println("This node already assigned");
		return false
	}

	parentNode.right = rightNode
	
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

func isFullBinaryTree(tree *Node) bool {
	if tree == nil {
		return true
	}

	if tree.left == nil && tree.right == nil {
		return true
	}

	if tree.left != nil && tree.right != nil {
		return isFullBinaryTree(tree.left) && isFullBinaryTree(tree.right)
	}

	return false
}

func main() {
	root := createNode(1)
	node2 := createNode(2)
	node3 := createNode(3)
	node4 := createNode(4)
	node5 := createNode(5)

	insertLeft(root, node2)
	insertRight(root, node3)

	insertLeft(node2, node4)
	insertRight(node2, node5)

	fmt.Print("Preorder Traversal")
	preorderTraversal(root)

	fmt.Print("\nInorder Traversal")
	inorderTraversal(root)

	fmt.Print("\nPostorder Traversal")
	postorderTraversal(root)

	fmt.Println("\nIs Full Binary Tree:", isFullBinaryTree(root))
}