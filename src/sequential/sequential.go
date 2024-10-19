package sequential

import (
	"fmt"
	// @ts-ignore
	utilities "lab3/src/utils"
	"strconv"
	"strings"
	"time"
)

func BSTSeqential(data []string, args *utilities.ArgumentParser){

	// Store the start time
	timer := utilities.Timer{Start: time.Now()}

	// Map of hash numbers to slice of root nodes
	hashes := make(map[int][]*utilities.BSTRootNode)

	for ID, line := range data {

		// Create an empty slice of ints to hold the BST values read in
		bst := make([]int, 0)

		// Iterate read in line split by " " and assign int converted value to bst slice
		for _, value := range strings.Split(line, " ") {
			val, err := strconv.Atoi(value)
			if err == nil {
				bst = append(bst, val)
			}

		}

		// Build the BST and calculate Hash # and assign to Hash map
		if len(bst) > 1 {
			root := BuildBST(bst, ID) // Returns the root of the BST
			root.GenHashNumber(root.Root, args) // Generate the Hash Value and In place order
			
			 // Print a new line because GenHashNumbers prints out in place order
			if *args.IsPrint{fmt.Printf("\n")}

			// Append the bst to a Hash group or create a new group if non-existent
			hashes[root.Hash] = append(hashes[root.Hash], root)
		}
	}

	// Print sequential Hashtime
	fmt.Printf("hashTime: %f\n", timer.TrackTime().Seconds())
	
	// Print Hash groups
	if *args.Data_workers > 0 {
		utilities.PrintHashGroups(timer.TrackTime().Seconds(), hashes)
	}
	
	// Print Comp groups
	if *args.Comp_workers > 0 {
		// CompareBST()
	}

}

func BuildBST(data []int, id int) *utilities.BSTRootNode {
	root := utilities.BSTRootNode{ID: id, Root: &utilities.Node{Value: data[0], Left: nil, Right: nil}, Hash: 1, InPlaceOrder: ""}
	for i, value := range data {
		if i == 0 { 
			continue
		} else {
			root.InsertNode(value)
		}
	}

	return &root
}

// func CompareBST(hashes map[int][]*utilities.BSTRootNode) {
// 	groups := make(map[int][]*utilities.BSTRootNode)

// 	for hash, group := range hashes {

		

// 	}
// }