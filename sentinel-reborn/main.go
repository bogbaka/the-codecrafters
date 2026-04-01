package main

import (
	"fmt"
	"os"
)

func main() {

	// Ensure the program receives exactly two arguments: input and output files
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . input.txt output.txt")
		return
	}

	// Get input and output file paths
	inputFile := os.Args[1]
	outputFile := os.Args[2]

	// Read the content of the input file
	data, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	result := textprocessor.ProcessText(string(data))

	// Write the processed text to the output file
	err = os.WriteFile(outputFile, []byte(result+"\n"), 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
}
