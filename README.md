# Codecrafters

## Objectives

In this project we will use some of our old functions made in our code crafters works. we  used them with the objective of making a simple text completion/editing/auto-correction tool.

One more detail. This time the project will be corrected by Code Legends. The Code Legends are other students.

## Introduction

    Our project will be written in Go.

    The code should respect the good practices.
Objectives

In this project you will use some of your old functions made in your old repository. You will use them with the objective of making a simple text completion/editing/auto-correction tool.

One more detail. This time the project will be corrected by auditors. The auditors will be other students and you will be an auditor as well.

The tool we built will receive as arguments the name of a file containing a text that needs some modifications (the input) and the name of the file the modified text should be placed in (the output). Next is a list of possible modifications that Our program should execute:

## MODIFY FUNCTION
It was by (Jonathan Ahubi and Daniel Akpa)
The modify function is to go throuth the inputed string and do all this changes:
1. Convert hexadecimal to decimal
2. Convert binary to decimal
3. Any unwanted lowercase change to uppercase 
4. Any unwanted uppercase change to lowercase 
5. Capitalise

## Tokenizer
This package provides a simple tokenizer for splitting a string into meaningful tokens.
Done by (Raymond and Moses)

### Functions

- **Punctuation(c rune) bool**  
  Checks whether a character is a punctuation mark (`. , ; : ! ?`).

- **Tokenizer(tokens string) []string**  
  Splits an input string into tokens based on:
  - Spaces
  - Punctuation (grouped together)
  - Parentheses (keeps content inside `()` as a single token)
  - Apostrophes (`'`) treated as separate tokens

### Notes
- Text inside parentheses is preserved as one token.
- Consecutive punctuation marks are grouped into a single token.
- Words are separated cleanly without losing special characters.

## Processor Package

This package provides utilities for rebuilding text from a slice of tokens while properly handling punctuation and quotes.

### Function rebuild 
done by (daniel and ebo)

  Checks if a token is a single quote (`'`).
- 
- **Punctuation1(c string) bool**  
  Determines whether a token is a punctuation mark (`, . ? ! : ; '`).

- **Rebuild(tokens []string) string**  
  Reconstructs a string from tokens by:
  - Adding spaces between words where appropriate
  - Avoiding spaces before punctuation
  - Correctly handling opening and closing quotes using a toggle mechanism

## Article Fixer

This utility adjusts indefinite articles in a slice of tokens for better grammatical correctness.

### Function article

done by (regina)

- **FixArticle(token []string) []string**  
  Updates occurrences of `"a"` or `"A"` based on the following word:
  - Changes `"a"` → `"an"` when the next word starts with a vowel or `h`
  - Preserves capitalization for uppercase `"A"`



## Main Program
done by (blessing)

This is the entry point of the application. It processes a text file and outputs a cleaned, formatted version.

### Usage

```bash
go run . input.txt output.txt