package main

import (
	"fmt"
	sequential "lab3/src/sequential"
	utilities "lab3/src/utils"
)
func main() {

	hash_workers, data_workers, comp_workers, input_file, isPrint := utilities.ArgumentParser()

	if isPrint{
		fmt.Printf("CMD Args:\nInput: %s\nHash Workers: %d\nData Workers: %d\nComp Workers: %d\n" ,input_file, hash_workers, data_workers, comp_workers)
	}

	lines := utilities.FileReader(input_file, isPrint)

	if isPrint{
			for _, item := range lines {
		fmt.Println(item)
	}
	}

	if hash_workers == 1 && data_workers == 1 {
		sequential.BSTSeqential(lines, isPrint)
	}
}