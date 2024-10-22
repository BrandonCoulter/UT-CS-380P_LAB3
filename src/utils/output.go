package utilities

import (
	"fmt"
)

func PrintHashGroups(time float64, hashes map[int][]*BSTRootNode){
	fmt.Printf("hashGroupTime: %f Num Trees: %d\n", time, len(hashes))
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

func PrintCompTree(time float64, comps map[string][]*BSTRootNode){
	fmt.Printf("compareTreeTime: %f\n", time)
	group_count := 0
	for comp, group := range comps {
		if len(comps[comp]) > 1 {
			fmt.Printf("group %d: ", group_count)
			for _, node := range group {
				fmt.Printf("%d ", node.ID)
			}
			fmt.Printf("\n")
			group_count += 1
		}
	}

}