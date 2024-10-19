package utilities

import (
	"fmt"
	"io/ioutil"
	"strings"
)
func FileReader(file_path string, isPrint bool)[]string{
	if isPrint {
		fmt.Printf("Reading File: %s\n", file_path)
	}
	// Read file and extract data as a single string
	data, err := ioutil.ReadFile(file_path)
	if err != nil {
		fmt.Printf("An error occured reading the file: %s\n\tError: %s\n", file_path, err)
	}

	// Split string into single lines based on newline character
	split_data := strings.Split(string(data), "\n")

	// Create an slice of strings for each line
	lines := make([]string, 0)

	// Append each line to the slice
	lines = append(lines, split_data...)

	return lines // return the slice
}