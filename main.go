package main

import (
	"fmt"
	sequential "lab3/src/sequential"
	utilities "lab3/src/utils"
)
func main() {

	args := &utilities.ArgumentParser{}
	args.ParseArgs()

	if *args.IsPrint {
		fmt.Printf("CMD Args:\nInput: %s\nHash Workers: %d\nData Workers: %d\nComp Workers: %d\n" , *args.Input_file, *args.Hash_workers, *args.Data_workers, *args.Comp_workers)
	}

	lines := utilities.FileReader(*args.Input_file, *args.IsPrint)

	if *args.IsPrint{
		for _, item := range lines {
			fmt.Println(item)
		}
	}

	if *args.Hash_workers == 1 {
		sequential.BSTSeqential(lines, args)
	} else if *args.Hash_workers > 1 && *args.Data_workers == 1 {
		//Part 2
	} else if *args.Hash_workers > 1 && *args.Data_workers > 1 && *args.Hash_workers == *args.Data_workers {
		//Part 3
	} else if *args.Hash_workers > 1 && *args.Data_workers > 1 && *args.Hash_workers != *args.Data_workers {
		//Extra Credit
	}
}