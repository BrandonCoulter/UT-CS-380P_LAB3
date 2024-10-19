package utilities

import "fmt"


type BSTRootNode struct {
	ID int
	Root *Node
	Hash int
}

// TODO: I may not need this function, check back after testing
// Create the root node of a BST 
// func (bst *BSTRootNode) InsertRoot(value int) {
// 	bst.root = InsertNode(bst.root, value)
// 	bst.Hash = 1

// }

func (bst *BSTRootNode) InsertNode(value int) {
	InsertNode(bst.Root, value)
}

// Generate the hash number for a BST using in order traversal
func (bst *BSTRootNode) GenHashNumber(node *Node, isPrint bool){
	if node != nil {
		bst.GenHashNumber(node.Left, isPrint)
		if isPrint { fmt.Printf(" %d ", node.Value) }
		bst.Hash = AddToHash(bst.Hash, node.Value)
		bst.GenHashNumber(node.Right, isPrint)
	}
}

// Add to hash number based on function given in the instructions
func AddToHash(hash int, value int) int {
	new_value := value + 2

	hash = ((hash * new_value) + new_value) % 1000

	return hash
}

/**************************/
/*Binary Search Tree Nodes*/
/**************************/

type Node struct {
	Value int
	Left *Node
	Right *Node
}

func InsertNode(node *Node, value int) *Node {
	if node == nil {
		return &Node{Value: value}
	}

	if value < node.Value {
		node.Left = InsertNode(node.Left, value)
	} else if value > node.Value {
		node.Right = InsertNode(node.Right, value)
	}
	return node
}