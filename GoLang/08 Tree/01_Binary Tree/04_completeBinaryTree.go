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

// Count the number of nodes
func countNumNodes(root *Node) int {
	if root == nil {
		return 0
	}

	return 1 + countNumNodes(root.left) + countNumNodes(root.right)
}


// Check if the tree is a complete binary tree
func isCompleteBinaryTree(root *Node) bool {
	if root == nil {
		return true
	}

	nodeCount := countNumNodes(root)
	fmt.Println("\nNode Count: ", nodeCount)

	return checkComplete(root, nodeCount, 0)
}

func checkComplete(root *Node, numberNodes, index int) bool {
	// Check if the tree is complete
	if root == nil {
		return true
	}

	if index >= numberNodes {
		return false
	}

	return checkComplete(root.left, numberNodes, 2*index+1) && checkComplete(root.right, numberNodes, 2*index+2)
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

	if isCompleteBinaryTree(root) {
		fmt.Println("\nThe tree is a complete binary tree.")
	} else {
		fmt.Println("\nThe tree is not a complete binary tree.")
	}
}