package utilities

import (
	"fmt"
)

func PrintHashGroups(time float32, hashes []interface{}){
	fmt.Printf("hashGroupTime: %f\n", time)
	for _, hash := range hashes {
		fmt.Println(hash) //FIXME: Change the print hash output to match specifications
	}
}