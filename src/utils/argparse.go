package utilities

import (
	"flag"
)

type ArgumentParser struct {
	Hash_workers *int
	Data_workers *int
	Comp_workers *int
	Input_file *string
	IsPrint *bool
	
}

func (args *ArgumentParser)ParseArgs(){
	
	// Lab 3 required flags
	args.Hash_workers = flag.Int("hash-workers", 1, "integer-valued number of threads")
	args.Data_workers = flag.Int("data-workers", -1, "integer-valued number of threads")
	args.Comp_workers = flag.Int("comp-workers", -1, "integer-valued number of threads")
	args.Input_file = flag.String("input", "", "string-valued path to an input file")

	// Custom flags
	args.IsPrint = flag.Bool("p", false, "Enable debug printouts")
	
	flag.Parse()
} 
