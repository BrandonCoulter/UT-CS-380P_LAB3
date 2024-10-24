package parallel

import (
    utilities "lab3/src/utils"
)

// Struct to hold messages used by the channel
type Message struct {
    Node *utilities.BSTRootNode
    Hash int
}

func CentralHashManager(update chan Message, done chan bool) *map[int][]*utilities.BSTRootNode {

    // Define the central map for hashes to root nodes
    hash_map := make(map[int][]*utilities.BSTRootNode)

    for {
        select {
        case msg := <- update:
            hash_map[msg.Hash] = append(hash_map[msg.Hash], msg.Node)
        case <- done:
            return &hash_map
        }
    }

    return &hash_map
}