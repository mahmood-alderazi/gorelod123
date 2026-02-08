package main

import "strings"

func handleCase(text string) string {

	words := strings.Split(text, " ")

	for i, char := range words {
		if char == "(up)" && i > 0 {
			words[i] = ""
			words[i-1] = strings.ToUpper(words[i-1])

		} else if char == "(low)" && i > 0 {
			words[i] = ""
			words[i-1] = strings.ToLower(words[i-1])

		} else if char == "(cap)" && i > 0 {
			words[i] = ""
			words[i-1] = capitalize(words[i-1])
		}
	}
	return strings.Join(words, " ")
}

func capitalize(word string) string {
	if len(word) == 0 {
		return word
	}
	return strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
}
