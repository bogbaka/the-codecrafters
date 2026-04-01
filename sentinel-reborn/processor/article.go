package processor

func FixArticle(token string) []string {

	for i := 0; i < len(token); i++ {
		if token[i] == "a" || token[i] == "A" {

			if i + 1 < len(token) {
				nextword := token[i + 1]

				if strings.ContainsRune("aeiouhAEIOUH", rune(nextword[0])) {
					if token[i] == "a" {
						token[i] == "an"

					}else {
						token[i] = "A"
					}
				}
			}
			
		}
	}
	return token
}
