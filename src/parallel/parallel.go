package parallel

// import (
// 	"fmt"
// 	// @ts-ignore
// 	utilities "lab3/src/utils"
// 	"strconv"
// 	"strings"
// 	"sync"
// 	"time"
// )

// func BSTParallel(data []string, args *utilities.ArgumentParser) {
//     // Store the start time
// 	timer := utilities.Timer{Start: time.Now()}
    
//     // Create a list of binary Search trees
//     bst_trees := make([]*utilities.BSTRootNode, 0)
//     // Map of hash numbers to slice of root nodes
//     hashes := make(map[int][]*utilities.BSTRootNode)
//     // Create an empty map to hold each group
//     groups := make(map[string][]*utilities.BSTRootNode)

//     // Channel for passing node to get returned
//     node_channel := make(chan *utilities.BSTRootNode)
//     var wait_group sync.WaitGroup

//     // Iterate lines in the input file and Run a GO Routine for every BST
// 	for ID, line := range data {
//         wait_group.Add(1)
//         go BuildBST(ID, line, node_channel, &wait_group, args)
    
//     }

//     // Start GO Routine to collect all the root nodes (Central Manager)
//     go func() {
//         wait_group.Wait()    
//         close(node_channel)
//     }()
    
//     for root := range node_channel{
//             bst_trees = append(bst_trees, root)
//     }
    

//     // Print sequential Hashtime
// 	fmt.Printf("hashTime: %f Num Trees: %d\n", timer.TrackTime().Seconds(), len(bst_trees))
    
// 	//Find hash groups and print them
// 	if *args.Data_workers > 0 {

//         timer.Start = time.Now() // Restart timer for Hash BST calculation/grouping
        
//         // Channel for passing hashes to get mapped
//         hash_channel := make(chan *utilities.BSTRootNode)

//         for _, root := range bst_trees {
//             wait_group.Add(1)
//             go GroupBST(root, hash_channel, &wait_group, args)
//         }

//             // Start GO Routine to collect all the root nodes (Central Manager)
//         go func() {
//             wait_group.Wait()    
//             close(hash_channel)
//         }()

//         for root := range hash_channel{
//             // Append the bst to a Hash group or create a new group if non-existent
//             hashes[root.Hash] = append(hashes[root.Hash], root)
//         }
        
//         utilities.PrintHashGroups(timer.TrackTime().Seconds(), hashes) // Print the hash groups
// 	}

//     if *args.Comp_workers > 0 {

//         timer.Start = time.Now() // Restart timer for Compare BST calculation/grouping

//         // Channel for passing hashes to get mapped
//         comp_channel := make(chan *utilities.CompareResult)

//         // For each hash group in the hashes check if they are 
//         for _, group := range hashes {
//             // If there are more then 1 BST in the hash group
//             // If there is only 1, there need not be a reason to check it
//             if len(group) > 1 {
//                 for _, node := range group {
//                     wait_group.Add(1)
//                     go CompareBST(node, comp_channel, &wait_group)
//                 }
//             }
//         }
//         go func() {
//             wait_group.Wait()    
//             close(comp_channel)
//         }()

//         for result := range comp_channel{
//             // Append the bst to a Hash group or create a new group if non-existent
//             groups[result.GroupID] = append(groups[result.GroupID], result.Node)
//         }

//         utilities.PrintCompTree(timer.TrackTime().Seconds(), groups) // Print groups
//     }

// }

// func BuildBST(id int, line string, nc chan *utilities.BSTRootNode, wg *sync.WaitGroup, args *utilities.ArgumentParser) {

//     // Ensure that the wait_group.Done() method is called regardless of return status
//     defer wg.Done()

//     // Create an empty slice of ints to hold the BST values read in
//     bst := make([]int, 0)

//     // Iteratively read in line split by " " and assign int converted value to bst slice
//     for _, value := range strings.Split(line, " ") {
//         val, err := strconv.Atoi(value)
//         if err == nil {
//             bst = append(bst, val)
//         }
//     }

//     // Build the BST and calculate Hash #
//     if len(bst) > 0 {
//         // Create the root node with ID, and intial Hash set to 1
//         root := utilities.BSTRootNode{ID: id, Root: &utilities.Node{Value: bst[0], Left: nil, Right: nil}, Hash: 1}

//         // Iterate values and assign them to nodes
//         for i, value := range bst {
//             // Skip value in index 0 since this is the root node
//             if i == 0 { 
//                 continue
//             } else {
//                 // Insert a node with the given value
//                 root.InsertNode(value)
//             }
//         }
        
//         // Generate the Hash Value without mapping In place order
//         root.GenHashNumber(root.Root, true, false)
//         // Appended the root node to the list of BSTs via the node_channel
//         nc <- &root
//     }

// }


// func GroupBST(root_node *utilities.BSTRootNode, hc chan *utilities.BSTRootNode, wg *sync.WaitGroup, args *utilities.ArgumentParser) {
    
//     // Ensure that the wait_group.Done() method is called regardless of return status
//     defer wg.Done()

//     // Regenerate the Hash value with mapping In place order 
//     root_node.GenHashNumber(root_node.Root, false, true)
//     // Print a new line because GenHashNumbers prints out in place order
//     if *args.IsPrint{fmt.Printf("\n")}

//     hc <- root_node // Return root with mapped In place order value slice via channel
// }

// func CompareBST(root_node *utilities.BSTRootNode, cc chan *utilities.CompareResult, wg *sync.WaitGroup) {

//     defer wg.Done()
    
//     result := utilities.CompareResult{GroupID: fmt.Sprint(root_node.InPlaceOrder), Node: root_node}

// 	// Return the group info
//     cc <- &result
// }