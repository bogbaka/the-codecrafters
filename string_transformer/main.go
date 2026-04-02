package main

import (
	//"bufio"
	"bufio"
	"fmt"
	"os"

	//"os"
	"strings"
)

func processLine(line string, lineNumber int) (string, bool) {

	// 1. Trim
	line = strings.TrimSpace(line)

	// 2. replace TODO
	if strings.HasPrefix(line, "TODO") {
		line = strings.Replace(line, "TODO:", "✦ ACTION:", 1)
	}

	// 3. Replace CLASSIFIED
	if strings.HasPrefix(line, "CLASSIFIED") {
		line = strings.Replace(line, "CLASSIFIED", "[REDACTED]:", 1)
	}

	// 4. Remove empty or dash lines
	if line == "" || line == "-----------------------------------------------" {
		return "", false
	}

	// 5. Add line number
	numbered := fmt.Sprintf("%03d. %s", lineNumber, line)

	return numbered, true
}

func main() {

	// Check Arguments
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . <input.txt> <output.txt>")
		return
	}
	inputFile := os.Args[1]
	outputFile := os.Args[2]

	// Prevent Same File
	if inputFile == outputFile {
		fmt.Println("✗ File not found:", inputFile)
		return
	}
	// OPEN INPUT FILE
	inFile, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("✗ File not found:", inputFile)
		return
	}
	defer inFile.Close()

	// create output file
	outFile, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("✗ Cannot write to output file")
		return
	}
	defer inFile.Close()

	scanner := bufio.NewScanner(inFile)

	lineNumber := 1
	linesRead := 0
	linesWritten := 0
	linesRemoved := 0

	for scanner.Scan() {
		line := scanner.Text()
		linesRead++

		result, keep := processLine(line, lineNumber)

		if keep {
			outFile.WriteString(result + "\n")
			lineNumber++
			linesWritten++
		} else {
			linesRemoved++
		}
	}
	// Handle empty File
	if linesRead == 0 {
		fmt.Println("⚠ Input file is empty. Nothing to process.")
	}
	// Print Summary
	fmt.Println("✦ Lines read    :", linesRead)
	fmt.Println("✦ Lines written :", linesWritten)
	fmt.Println("✦ Lines removed :", linesRemoved)
	fmt.Println("✦ Rules applied : Trim, TODO→ACTION, CLASSIFIED→REDACTED, Remove empty/dash, Add numbering")
}
