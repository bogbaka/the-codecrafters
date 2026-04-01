package processor

import "strings"

// helper function to check single Quote
func isQuote(token string) bool {
	return token == "'"
}

// helper function to check punctuation
func Punctuation1(c string) bool {
	switch c {
	case ",", ".", "?", "!", ":", ";", "'":
		return true
	}
	return false

}

// Rebuild take the slice of string (tokens), fix Quote and punctuation and return a string
func Rebuild(tokens []string) string {

	// using toggle logic to handle quote
	isInsideQoute := false

	// b as a builder to build string
	var b strings.Builder

	for i := 0; i < len(tokens); i++ {
		token := tokens[i]

		// handle "space before" logic

		if i-1 > 0 {

			prev := tokens[i-1]
			// Add space if not pauntuation
			// if not inside the openig and before closing quote

			isClosingQoute := token == "'" && isInsideQoute

			isopeningQoutePrev := prev == "'" && isInsideQoute

			if !Punctuation1(token) && !isClosingQoute && !isopeningQoutePrev {
				b.WriteRune(' ')
			}

			if isQuote(tokens[i]) && !isInsideQoute {
				b.WriteRune(' ')
			}
		}
		// defualt write string
		b.WriteString(tokens[i])

		// toggle if seen quote
		if isQuote(token) {
			isInsideQoute = !isInsideQoute
		}

	}
	return b.String()
}
