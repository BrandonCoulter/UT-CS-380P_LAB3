package sequential

import (
	"fmt"
	// @ts-ignore
	utilities "lab3/src/utils"
	"strconv"
	"strings"
	"time"
)

// BST Seqential implementation
func BSTSeqential(data []string, args *utilities.ArgumentParser){

	// Store the start time
	timer := utilities.Timer{Start: time.Now()}

    // Create a list of binary Search trees
    bst_trees := make([]*utilities.BSTRootNode, 0)

    // Map of hash numbers to slice of root nodes
	hashes := make(map[int][]*utilities.BSTRootNode)

	// Iterate lines in the input file
	for ID, line := range data {

		// Create an empty slice of ints to hold the BST values read in
		bst := make([]int, 0)

		// Iteratively read in line split by " " and assign int converted value to bst slice
		for _, value := range strings.Split(line, " ") {
			val, err := strconv.Atoi(value)
			if err == nil {
				bst = append(bst, val)
			}

		}

		// Build the BST and calculate Hash #
		if len(bst) > 0 {
			root := BuildBST(bst, ID) // Returns the root of the BST
			root.GenHashNumber(root.Root, args, false) // Generate the Hash Value without In place order

            bst_trees = append(bst_trees, root) // Appended the root node to the list of BSTs
		}
	}

	// Print sequential Hashtime
	fmt.Printf("hashTime: %f\n", timer.TrackTime().Seconds())
	
	// Print find hash groups and preint them
	if *args.Data_workers > 0 {
        timer.Start = time.Now() // Restart timer for Hash BST calculation/grouping
        hashes = GroupBST(bst_trees, args)

        utilities.PrintHashGroups(timer.TrackTime().Seconds(), hashes) // Print the hash groups
        
	}
	
	// Print Comp groups
	if *args.Comp_workers > 0 {
		timer.Start = time.Now() // Restart timer for Compare BST calculation/grouping
		groups := CompareBST(hashes) // Collect groups
		utilities.PrintCompTree(timer.TrackTime().Seconds(), groups) // Print groups
	}

}

// Function to build the BST
func BuildBST(data []int, id int) *utilities.BSTRootNode {
	// Create the root node with ID, and intial Hash set to 1
	root := utilities.BSTRootNode{ID: id, Root: &utilities.Node{Value: data[0], Left: nil, Right: nil}, Hash: 1}
	// Iterate values and assign them to nodes
	for i, value := range data {
		// Skip value in index 0 since this is the root node
		if i == 0 { 
			continue
		} else {
			// Insert a node with the given value
			root.InsertNode(value)
		}
	}

	// Return a reference to the root node
	return &root
}

func GroupBST(bst_trees []*utilities.BSTRootNode, args *utilities.ArgumentParser) map[int][]*utilities.BSTRootNode {

    // Map of hash numbers to slice of root nodes
	hashes := make(map[int][]*utilities.BSTRootNode)
    
    for _, root := range(bst_trees) {
        
        root.GenHashNumber(root.Root, args, true) // Regenerate the Hash value with In place order 
        // Print a new line because GenHashNumbers prints out in place order
        if *args.IsPrint{fmt.Printf("\n")}

        // Append the bst to a Hash group or create a new group if non-existent
        hashes[root.Hash] = append(hashes[root.Hash], root)
    }

    return hashes
}

// Function to compare BST based on in-order traversal hash
func CompareBST(hashes map[int][]*utilities.BSTRootNode) map[string][]*utilities.BSTRootNode {
	
	// Create an empty map to hold each group
	groups := make(map[string][]*utilities.BSTRootNode)

	// For each hash group in the hashes check if they are 
	for _, group := range hashes {
		// If there are more then 1 BST in the hash group
		// If there is only 1, there need not be a reason to check it
		if len(group) > 1 {
			for _, node := range group {
				groupID := fmt.Sprint(node.InPlaceOrder) // Create a in-order traversal key
				groups[groupID] = append(groups[groupID], node) // Map that key to each node it applys too
			}
		}
	}
	// Return the groups
	return groups
}