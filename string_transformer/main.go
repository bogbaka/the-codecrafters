package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// UPPER
func capEachWord(word string) string {
	return strings.ToUpper(word)
}

// TITLE CASE
func Title(text string) string {
	smallWords := map[string]bool{
		"a": true, "an": true, "the": true,
		"and": true, "but": true, "or": true,
		"for": true, "nor": true,
		"on": true, "at": true, "to": true, "by": true,
		"in": true, "of": true, "up": true,
		"as": true, "is": true, "it": true,
	}

	words := strings.Fields(text)
	result := ""

	for i := 0; i < len(words); i++ {
		word := strings.ToLower(words[i])

		if i == 0 || !smallWords[word] {
			result += strings.ToUpper(word[:1]) + word[1:]
		} else {
			result += word
		}

		if i != len(words)-1 {
			result += " "
		}
	}

	return result
}

// LOWER
func lower(word string) string {
	return strings.ToLower(word)
}

// CAP
func capWords(text string) string {
	words := strings.Fields(text)
	result := ""

	for i := 0; i < len(words); i++ {
		word := strings.ToLower(words[i])
		result += strings.ToUpper(word[:1]) + word[1:]

		if i != len(words)-1 {
			result += " "
		}
	}

	return result
}

// SNAKE
func Snake(text string) string {
	text = strings.ToLower(text)
	result := ""

	for i := 0; i < len(text); i++ {
		char := text[i]

		if (char >= 'a' && char <= 'z') || (char >= '0' && char <= '9') {
			result += string(char)
		} else if char == ' ' {
			result += "_"
		}
	}

	return result
}

// REVERSE
func Reverse(text string) string {
	words := strings.Fields(text)

	for i, w := range words {
		runes := []rune(w)

		for a, b := 0, len(runes)-1; a < b; a, b = a+1, b-1 {
			runes[a], runes[b] = runes[b], runes[a]
		}

		words[i] = string(runes)
	}

	return strings.Join(words, " ")
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	for {

		fmt.Print("Enter command: ")

		if !scanner.Scan() {
			break
		}

		input := scanner.Text()

		word := strings.Fields(input)
		command := strings.Join(word[:1], " ")

		text := strings.Join(word[1:], " ")

		// fmt.Scan(&command, &text)

		command = strings.ToLower(strings.TrimSpace(command))
		text = strings.TrimSpace(text)

		// EXIT
		if command == "exit" {
			fmt.Println("Shutting down String Transformer. Goodbye.")
			break
		}

		// MISSING TEXT
		if text == "" {
			fmt.Println("✗ No text provided. Usage: command <text>")
			continue
		}

		// COMMAND HANDLER
		switch command {

		case "upper":
			fmt.Println(strings.ToUpper(text))

		case "lower":
			fmt.Println(lower(text))

		case "cap":
			fmt.Println(capWords(text))

		case "title":
			fmt.Println(Title(text))

		case "snake":
			fmt.Println(Snake(text))

		case "reverse":
			fmt.Println(Reverse(text))

		default:
			fmt.Println("✗ Unknown command:", command)
			fmt.Println("Valid commands: upper, lower, cap, title, snake, reverse, exit")
		}
	}
	// fmt.Println(capEachWord("sentinel is online"))
	// fmt.Println(lower("ALERT LEVEL FIVE DETECTED"))
	// fmt.Println(capEachWord("director adaeze okonkwo"))
	// fmt.Println(capEachWord("THREAT LEVEL elevated"))
	// fmt.Println(Reverse("Lagos Nigeria"))
	// fmt.Println(Snake("Operation Gopher Protocol"))
	// fmt.Println(Snake("Alert! Level 5 detected"))
	// fmt.Println(Title("the fall of the western power grid" ))
	// fmt.Println(Title("a threat in the north" ))

}
