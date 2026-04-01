package processor

import (
	"strconv"
	"strings"
	"unicode"
)

func Modifier(token []string) []string {
	for i := 0; i < len(token); i++ {
		words := token[i]
		if strings.HasPrefix(words, "(") && strings.HasSuffix(words, ")") {
			words_content := words[1 : len(words)-1]

			words_part := strings.Split(words_content, ",")

			command := words_part[0]
			count := 1

			if len(words_part) > 1 {
				count, _ = strconv.Atoi(strings.TrimSpace(words_part[1]))
			}

			switch command {
			case "hex":
				if i > 0 {
					number, err := strconv.ParseInt(token[i-1], 16, 64)
					if err != nil {
						continue
					}
					token[i-1] = strconv.FormatInt(number, 10)

				}
			case "bin":
				if i > 0 {
					number, err := strconv.ParseInt(token[i-1], 2, 64)
					if err != nil {
						continue
					}
					token[i-1] = strconv.FormatInt(number, 10)
				}
			case "up":
				start := i - count
				if start < 0 {
					start = 0
				}

				for j := start; j < i; j++ {
					token[j] = strings.ToUpper(token[j])
				}

			case "low":
				start := i - count
				if start < 0 {
					start = 0
				}
				for j := start; j < i; j++ {
					token[j] = strings.ToLower(token[j])
				}

			case "cap":
				start := i - count
				for j := start; j < i; j++ {
					word := strings.ToLower(token[j])
					runes := []rune(word)

					if len(runes) > 0 {
						runes[0] = unicode.ToUpper(runes[0])
						token[j] = string(runes)
					}

				}

			}
			token = append(token[:i], token[i+1:]...)
			i--

		}
	}
	return token
}
