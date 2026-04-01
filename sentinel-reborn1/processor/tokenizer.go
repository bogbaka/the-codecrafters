package processor

import "strings"

func Punctuation(c rune) bool {
	switch c {
	case '.', ',', ';', ':', '!', '?':
		return true
	}
	return false

}

func Tokenizer(tokens string) []string {
	var b strings.Builder
	runes := []rune(tokens)
	var result []string
	//
	insideBracket := false
	for i := 0; i < len(runes); i++ {
		char := runes[i]
		switch {
		case char == ')':
			insideBracket = false
			b.WriteRune(')')
			result = append(result, b.String())
			b.Reset()
			continue
		case char == '(':
			if b.Len() > 0 {
				result = append(result, b.String())
				b.Reset()
			}
			insideBracket = true
			b.WriteRune('(')
			continue
		case insideBracket:
			b.WriteRune(char)
			continue
		case char == ' ':
			if b.Len() > 0 {
				result = append(result, b.String())
				b.Reset()
			}
			continue

		case char == '\'':
			if b.Len() > 0 {
				result = append(result, b.String())
				b.Reset()
			}
			result = append(result, "'")
			continue
		case Punctuation(char):
			if b.Len() > 0 {
				result = append(result, b.String())
				b.Reset()
			}
			start := i
			for i+1 < len(runes) && Punctuation(runes[i+1]) {
				i++
			}
			result = append(result, string(runes[start:i+1]))
			continue
		default:
			b.WriteString(string(char))
		}
	}
	if b.Len() > 0 {
		result = append(result, b.String())
	}
	return result
}
