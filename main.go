package main

import (
	"fmt"
	parallel "lab3/src/parallel"
	sequential "lab3/src/sequential"
	utilities "lab3/src/utils"
)

// Main GO function for LAB #3
func main() {

	// Collect and parse CMD args
	args := &utilities.ArgumentParser{}
	args.ParseArgs()

	// Debug print statement
	if *args.IsPrint {
		fmt.Printf("CMD Args:\nInput: %s\nHash Workers: %d\nData Workers: %d\nComp Workers: %d\n" , *args.Input_file, *args.Hash_workers, *args.Data_workers, *args.Comp_workers)
	}

	// Read all the lines of the file
	lines := utilities.FileReader(*args.Input_file, *args.IsPrint)

	// Debug print statement
	if *args.IsPrint{
		for _, item := range lines {
			fmt.Println(item)
		}
	}

	// Run differnt test cases based on CMD arguments
	if *args.Hash_workers == 1 {
		sequential.BSTSeqential(lines, args) // Run seqential
	} else if *args.Hash_workers > 1 {
        parallel.BSTParallel(lines, args)
		//Part 2
	} else if *args.Hash_workers > 1 && *args.Data_workers > 1 && *args.Hash_workers == *args.Data_workers {
		//Part 3
	} else if *args.Hash_workers > 1 && *args.Data_workers > 1 && *args.Hash_workers != *args.Data_workers {
		//Extra Credit
	}
}