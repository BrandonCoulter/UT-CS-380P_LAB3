package utilities

import (
	"fmt"
)

//FIXME: Make sure thqt hash print out is correctly formated
func PrintHashGroups(time float64, hashes map[int][]*BSTRootNode){
	fmt.Printf("hashGroupTime: %f\n", time)
	for hash, group := range hashes {
		if len(hashes[hash]) > 1 {
			fmt.Printf("%d: ", hash)
			for _, node := range group {
				fmt.Printf("%d ", node.ID)
			}
			fmt.Printf("\n")
		}
	}
}

func PrintCompTree(time float64, comps map[int][]*BSTRootNode){
	fmt.Printf("compareTreeTime: %f\n", time)

}