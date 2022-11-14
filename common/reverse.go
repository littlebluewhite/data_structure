package common

func reverseCharacters[T any](arr *[]T, startIndex, endIndex int) {
	for startIndex < endIndex {
		(*arr)[startIndex], (*arr)[endIndex] = (*arr)[endIndex], (*arr)[startIndex]
		startIndex++
		endIndex--
	}
}

func ReverseWords(words []rune) []rune {
	startIndex := 0
	reverseCharacters(&words, 0, len(words)-1)
	for i, w := range words {
		if w == ' ' {
			reverseCharacters[rune](&words, startIndex, i-1)
			startIndex = i + 1
		}
	}
	reverseCharacters[rune](&words, startIndex, len(words)-1)
	return words
}
