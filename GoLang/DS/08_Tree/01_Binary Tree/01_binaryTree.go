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

	InsertLeft(node3, node6)

	fmt.Print("DisplayPreorder")
	DisplayPreorder(root)

	fmt.Print("\nDisplayInorder")
	DisplayInorder(root)

	fmt.Print("\nDisplayPostorder")
	DisplayPostorder(root)
}