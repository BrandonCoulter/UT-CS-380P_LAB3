package utilities


type BSTRootNode struct {
	Root *Node
	ID int
	Hash int
	InPlaceOrder []int
    Key string
}

func (bst *BSTRootNode) InsertNode(value int) {
	InsertNode(bst.Root, value)
}

// Generate the hash number for a BST using in order traversal
func (bst *BSTRootNode) GenHashNumber(node *Node, perform_hash bool, gen_in_place bool){
	if node != nil {
		bst.GenHashNumber(node.Left, perform_hash, gen_in_place)

		if perform_hash { bst.Hash = AddToHash(bst.Hash, node.Value) }
        if gen_in_place { bst.InPlaceOrder = append(bst.InPlaceOrder, node.Value) }
		
		bst.GenHashNumber(node.Right, perform_hash, gen_in_place)
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
	} else if value >= node.Value {
		node.Right = InsertNode(node.Right, value)
	}
	return node
}


type CompareResult struct {
	Node *BSTRootNode
	GroupID string
}