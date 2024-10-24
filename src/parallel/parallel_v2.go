package parallel

import (
	"fmt"
	utilities "lab3/src/utils"
	"strconv"
	"strings"
	"time"
)

func BSTParallel(data []string, args *utilities.ArgumentParser) {

    // Slice of Root node
    bst_tree := make([]*utilities.BSTRootNode, 0)

	// Map of hash numbers to slice of root nodes
	var hashes map[int][]*utilities.BSTRootNode

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
	*      CHANNELS      * 
	**********************/   

    // Hash Group Time
    update := make(chan Message)
    hash_grouped := make(chan bool)
    all_groupers_complete := make(chan bool)
    hash_group_complete := make(chan map[int][]*utilities.BSTRootNode)

	/*********************
	*     BST HASHING    * 
	**********************/

	// Store the start time
	timer := utilities.Timer{Start: time.Now()}

    // Spawn a hash worker for every BST
	GoPerBST(bst_tree)

	// Print sequential Hashtime
	fmt.Printf("hashTime: %f\n", timer.TrackTime().Seconds())

    /*********************
	*   HASH GROUPING    * 
	**********************/

	timer.Start = time.Now()

	go CentralHashManager(update, all_groupers_complete, hash_group_complete)

    // Spawn a hash worker for every BST
    for i := 0; i < len(bst_tree); i++ {
        go HashGroupWorker(bst_tree[i], update, hash_grouped)
    }

    // Collected all workers done messages before continuing
    for i := 0; i < len(bst_tree); i++ {
        <-hash_grouped
    }
    
    close(hash_grouped)

    all_groupers_complete <- true

    hashes = <- hash_group_complete

    close(update)
    close(hash_group_complete)

    // Print Hash groups
	if *args.Data_workers > 0 {
		utilities.PrintHashGroups(timer.TrackTime().Seconds(), hashes)
	}

}