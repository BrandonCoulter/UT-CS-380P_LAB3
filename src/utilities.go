package main

import (
	"flag"
	"fmt"
)

func ArgumentParser()(int, int, int, string, bool){
	
	// Lab 3 required flags
	hash_workers := flag.Int("hash-workers", -1, "integer-valued number of threads")
	data_workers := flag.Int("data-workers", -1, "integer-valued number of threads")
	comp_workers := flag.Int("comp-workers", -1, "integer-valued number of threads")
	input_file := flag.String("input", "", "string-valued path to an input file")

	// Custom flags
	isPrint := flag.Bool("print", false, "Enable debug printouts")
	
	flag.Parse()

	return *hash_workers, *data_workers, *comp_workers, *input_file, *isPrint
} 

func FileReader(file_path string, isPrint bool){
	if isPrint {
		fmt.Printf("Reading File: %s\n", file_path)
	}
}