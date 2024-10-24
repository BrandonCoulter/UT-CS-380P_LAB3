package parallel

import (
    utilities "lab3/src/utils"
)

func GoPerBST(bst_tree []*utilities.BSTRootNode){
    // Hash Time
    hash_complete := make(chan bool)
    
    // Spawn a hash worker for every BST
    for i := 0; i < len(bst_tree); i++ {
        go HashWorker(bst_tree[i], hash_complete)
    }

    // Collected all workers done messages before continuing
    for i := 0; i < len(bst_tree); i++ {
        <-hash_complete
    }
    close(hash_complete)

    return
}