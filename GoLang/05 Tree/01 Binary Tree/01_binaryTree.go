//? Binary tree
// A binary tree is a tree data structure in which each parent node can have at most two children. 
// Each node of a binary tree consists of three items:

//     data item
//     address of left child
//     address of right child


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

	insertLeft(node3, node6)

	fmt.Print("Preorder Traversal")
	preorderTraversal(root)

	fmt.Print("\nInorder Traversal")
	inorderTraversal(root)

	fmt.Print("\nPostorder Traversal")
	postorderTraversal(root)
}