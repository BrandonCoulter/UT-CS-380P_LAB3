package main

import (
	"fmt"
)
func main() {

	hash_workers, data_workers, comp_workers, input_file, isPrint := ArgumentParser()

	if isPrint{
		fmt.Printf("CMD Args:\nInput: %s\nHash Workers: %d\nData Workers: %d\nComp Workers: %d\n" ,input_file, hash_workers, data_workers, comp_workers)
	}

	FileReader(input_file, isPrint)
}