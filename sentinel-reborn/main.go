package main

import (
	"fmt"
	"os"
	"sentinel-reborn/processor"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Correct Usage go run . sample.txt result.txt")
		return
	}

	inputfile := os.Args[1]
	outputfile := os.Args[2]

	data, err := os.ReadFile(inputfile)
	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}

	text := string(data)
	tokens := processor.Tokenizer(text)
	tokens = processor.Modifier(tokens)
	tokens = processor.FixArticle(tokens)
	result := processor.Rebuild(tokens)

	if !strings.HasSuffix(result, "\n") {
		result += "\n"
	}
	err = os.WriteFile(outputfile, []byte(result), 0644)
	if err != nil {
		fmt.Println("Error Writing file", err)
		return
	}
	fmt.Println("Process completed")
}
