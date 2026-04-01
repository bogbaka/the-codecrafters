package main

import (
	"fmt"
	"mock-go-reloaded/processor"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Correct Usage go run . samplel.txt result.txt")
		return
	}

	inputfile := os.Args[1]
	outputfile := os.Args[2]

	data, err := os.ReadFile(inputfile)
	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}
}