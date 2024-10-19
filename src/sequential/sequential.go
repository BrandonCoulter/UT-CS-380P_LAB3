package sequential

import (
	"fmt"
	utilities "lab3/src/utils"
	"strconv"
	"strings"
	"time"
)

func BSTSeqential(data []string, isPrint bool){

	timer := utilities.Timer{Start: time.Now()}

	bsts := make([]*utilities.BSTRootNode, 0)

	for ID, line := range data {
		// fmt.Printf("Seq BST line: %s\n", line)

		bst := make([]int, 0)

		for _, value := range strings.Split(line, " ") {
			val, err := strconv.Atoi(value)
			if err == nil {
				bst = append(bst, val)
			}

		}
		if len(bst) > 1 {
			root := BuildBST(bst, ID)
			bsts = append(bsts, root)
		}
	}

	for _, bst := range bsts {
		if isPrint {
			fmt.Printf("\n")
		}
		bst.GenHashNumber(bst.Root, isPrint)
		fmt.Printf("ID: %d, %d\n", bst.ID, bst.Hash)
	}

	fmt.Printf("hashTime: %f\n", timer.TrackTime().Seconds())

}

func BuildBST(data []int, id int) *utilities.BSTRootNode {
	root := utilities.BSTRootNode{ID: id, Root: &utilities.Node{Value: data[0], Left: nil, Right: nil}, Hash: 1}
	for i, value := range data {
		if i == 0 { 
			continue
		} else {
			root.InsertNode(value)
		}
	}

	return &root
}