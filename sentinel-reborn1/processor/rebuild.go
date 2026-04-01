package processor

import(
	"fmt"
	"strings"
	
)


func Rebuild(tokens []string) string {
	isInsideQuote := false

	var b strings.Builder
	for i := 0; i < len(tokens); i++ {
		token := tokens[i]

		if i-1 > 0 {
			prev := tokens[i-1]

			isClosingQuote := token == "'" && isInsideQuote
			isOpeningQuotePrev := prev == "'" && isInsideQuote

		if !Punctuation1(token) && !isClosingQuote && !isOpeningQuotePrev {
				b.WriteRune(' ')
			}
		}
		b.WriteString(tokens[i])
		if isQuote(token) {
			isInsideQuote = !isInsideQuote
		}
	}
	return b.String()
}

