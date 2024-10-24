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

	bst_tree := make([]*utilities.BSTRootNode, 0)

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

		// Build the BST and calculate Hash # and assign to Hash map
		if len(bst) > 0 {
			root := utilities.BuildBST(bst, ID) // Returns the root of the BST
			root.GenHashNumber(root.Root, false, true) // Generate In place order without hash values
            root.Key = fmt.Sprintf("%v", root.InPlaceOrder)
			bst_tree = append(bst_tree, root)
		}
	}

	/*********************
	*     BST HASHING    * 
	**********************/

	// Store the start time
	timer := utilities.Timer{Start: time.Now()}

	HashBST(bst_tree)

	// Print sequential Hashtime
	fmt.Printf("hashTime: %f\n", timer.TrackTime().Seconds())

	/*********************
	*   HASH GROUPING    * 
	**********************/

	timer.Start = time.Now()

	hashes = GroupBST(bst_tree)

	// Print Hash groups
	if *args.Data_workers > 0 {
		utilities.PrintHashGroups(timer.TrackTime().Seconds(), hashes)
	}

	/*********************
	*    TREE COMPARE    * 
	**********************/

	timer.Start = time.Now() // Restart timer for Compare BST calculation/grouping
	groups := CompareBST(hashes) // Collect groups

	// Print Comp groups
	if *args.Comp_workers > 0 {
		utilities.PrintCompTree(timer.TrackTime().Seconds(), groups) // Print groups
	}

}

func HashBST(bst_tree []*utilities.BSTRootNode){
	for _, root := range bst_tree {
		root.GenHashNumber(root.Root, true, false) // Generate the Hash Value and without In place order
	}
}

func GroupBST(bst_tree []*utilities.BSTRootNode) map[int][]*utilities.BSTRootNode {
	// Map of hash numbers to slice of root nodes
	hashes := make(map[int][]*utilities.BSTRootNode)
	
	for _, root := range bst_tree {

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
			for _, root := range group {
				groups[root.Key] = append(groups[root.Key], root) // Map that key to each node it applys too
			}
		}
	}
	// Return the groups
	return groups
}