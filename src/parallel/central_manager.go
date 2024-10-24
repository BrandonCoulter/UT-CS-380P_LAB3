package parallel

import (
	utilities "lab3/src/utils"
)

// Struct to hold messages used by the channel
type Message struct {
    Root *utilities.BSTRootNode
    Hash int
}

func CentralHashManager(update chan Message, all_groupers_complete chan bool, hash_group_complete chan map[int][]*utilities.BSTRootNode){

    // Define the central map for hashes to root nodes
    hash_map := make(map[int][]*utilities.BSTRootNode)

    for {
        select {
        case msg := <- update:
            hash_map[msg.Hash] = append(hash_map[msg.Hash], msg.Root)
        case <- all_groupers_complete:
            hash_group_complete <- hash_map
            return
        }
    }
}

func HashWorker(root *utilities.BSTRootNode, hashed chan bool){
    // Generate the Hash Value and without In place order
    root.GenHashNumber(root.Root, true, false)

    // Let the main thread know that this worker is completed
    hashed <- true
}

func HashGroupWorker(root *utilities.BSTRootNode, update chan Message, grouped_hash chan bool){
    // Formulate a message to send via update channel
    msg := Message{Root: root, Hash: root.Hash}

    // Update the central map
    update <- msg

    // Let the main thread know that this worker is completed
    grouped_hash <- true
}